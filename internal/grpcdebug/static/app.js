const state = {
  catalog: [],
  search: "",
  selectedMethod: "",
  methodDetail: null,
  detailCache: new Map(),
  mode: "form",
  formData: {},
  requestJson: "{}",
  result: null,
};

const elements = {
  catalog: document.getElementById("catalog"),
  searchInput: document.getElementById("searchInput"),
  targetInput: document.getElementById("targetInput"),
  timeoutInput: document.getElementById("timeoutInput"),
  authorityInput: document.getElementById("authorityInput"),
  serverNameInput: document.getElementById("serverNameInput"),
  useTlsInput: document.getElementById("useTlsInput"),
  skipVerifyInput: document.getElementById("skipVerifyInput"),
  metadataInput: document.getElementById("metadataInput"),
  methodPanel: document.getElementById("methodPanel"),
  requestPanel: document.getElementById("requestPanel"),
  invokeButton: document.getElementById("invokeButton"),
  invokeHint: document.getElementById("invokeHint"),
  resultPanel: document.getElementById("resultPanel"),
};

boot().catch((error) => {
  elements.methodPanel.classList.remove("empty-state");
  elements.methodPanel.innerHTML = renderFatal(`初始化失败: ${error.message}`);
});

async function boot() {
  bindEvents();
  await loadCatalog();
}

function bindEvents() {
  elements.searchInput.addEventListener("input", (event) => {
    state.search = event.target.value.trim().toLowerCase();
    renderCatalog();
  });

  elements.invokeButton.addEventListener("click", () => {
    invokeCurrentMethod().catch((error) => {
      renderResult({
        error: {
          code: "CLIENT",
          message: error.message,
          raw: error.stack || error.message,
        },
      });
    });
  });
}

async function loadCatalog() {
  const response = await fetch("/api/catalog");
  if (!response.ok) {
    throw new Error("无法加载接口目录");
  }

  const payload = await response.json();
  state.catalog = payload.services || [];
  renderCatalog();

  const firstMethod = state.catalog.flatMap((service) => service.methods || [])[0];
  if (firstMethod) {
    await selectMethod(firstMethod.fullMethod);
  }
}

function renderCatalog() {
  const fragment = document.createDocumentFragment();
  const keyword = state.search;

  const filtered = state.catalog.filter((service) => {
    if (!keyword) {
      return true;
    }
    const serviceMatches = service.fullName.toLowerCase().includes(keyword);
    const methodMatches = (service.methods || []).some((method) =>
      `${method.name} ${method.fullMethod}`.toLowerCase().includes(keyword),
    );
    return serviceMatches || methodMatches;
  });

  if (!filtered.length) {
    elements.catalog.innerHTML = `<div class="message-card"><p class="muted">没有匹配到接口。</p></div>`;
    return;
  }

  filtered.forEach((service) => {
    const detail = document.createElement("details");
    detail.open = true;

    const summary = document.createElement("summary");
    summary.innerHTML = `
      ${escapeHTML(service.fullName)}
      <span class="service-meta">${escapeHTML(service.file)} · ${(service.methods || []).length} methods</span>
    `;
    detail.appendChild(summary);

    const methods = document.createElement("div");
    methods.className = "method-list";

    (service.methods || [])
      .filter((method) => {
        if (!keyword) {
          return true;
        }
        return `${service.fullName} ${method.name} ${method.fullMethod}`.toLowerCase().includes(keyword);
      })
      .forEach((method) => {
        const button = document.createElement("button");
        button.type = "button";
        button.className = `method-item${method.fullMethod === state.selectedMethod ? " active" : ""}`;
        button.innerHTML = `
          <div>${escapeHTML(method.name)}</div>
          <div class="method-type">
            <span>${escapeHTML(shortType(method.requestType))}</span>
            <span>→</span>
            <span>${escapeHTML(shortType(method.responseType))}</span>
            ${method.clientStreaming || method.serverStreaming ? "<span class=\"badge accent\">stream</span>" : ""}
          </div>
        `;
        button.addEventListener("click", () => {
          selectMethod(method.fullMethod).catch((error) => {
            renderResult({
              error: {
                code: "CLIENT",
                message: error.message,
                raw: error.stack || error.message,
              },
            });
          });
        });
        methods.appendChild(button);
      });

    detail.appendChild(methods);
    fragment.appendChild(detail);
  });

  elements.catalog.innerHTML = "";
  elements.catalog.appendChild(fragment);
}

async function selectMethod(fullMethod) {
  state.selectedMethod = fullMethod;
  state.result = null;
  renderCatalog();
  renderResult(null);

  if (!state.detailCache.has(fullMethod)) {
    const response = await fetch(`/api/method?fullMethod=${encodeURIComponent(fullMethod)}`);
    if (!response.ok) {
      throw new Error(`加载方法详情失败: ${fullMethod}`);
    }
    state.detailCache.set(fullMethod, await response.json());
  }

  state.methodDetail = state.detailCache.get(fullMethod);
  state.formData = {};
  state.requestJson = "{}";
  state.mode = "form";

  renderMethod();
  renderRequest();
}

function renderMethod() {
  const detail = state.methodDetail;
  if (!detail) {
    elements.methodPanel.className = "panel empty-state";
    elements.methodPanel.innerHTML = `
      <div>
        <p class="eyebrow">方法详情</p>
        <h2>选择左侧接口后开始调试</h2>
      </div>
    `;
    return;
  }

  elements.methodPanel.className = "panel";
  elements.methodPanel.innerHTML = `
    <div class="method-summary">
      <div class="panel-head">
        <div>
          <p class="eyebrow">方法详情</p>
          <h2>${escapeHTML(detail.summary.fullMethod)}</h2>
          <p class="muted">${escapeHTML(detail.protoFile)}</p>
        </div>
      </div>
      <div class="pill-row">
        <span class="pill">Service: ${escapeHTML(detail.serviceName)}</span>
        <span class="pill">Request: ${escapeHTML(detail.summary.requestType)}</span>
        <span class="pill">Response: ${escapeHTML(detail.summary.responseType)}</span>
        ${detail.unary ? "<span class=\"pill\">Unary RPC</span>" : `<span class="pill warn">${escapeHTML(detail.streamingHint || "Streaming RPC")}</span>`}
      </div>
      ${
        detail.unary
          ? ""
          : `<p class="muted">当前页面先支持一元 RPC。这个方法是流式接口，暂时无法直接调试调用。</p>`
      }
    </div>
  `;
}

function renderRequest() {
  const detail = state.methodDetail;
  if (!detail) {
    elements.requestPanel.classList.add("hidden");
    elements.invokeButton.disabled = true;
    elements.invokeHint.textContent = "先从左侧选一个方法。";
    return;
  }

  elements.requestPanel.classList.remove("hidden");
  elements.invokeButton.disabled = !detail.unary;
  elements.invokeHint.textContent = detail.unary
    ? "调用前会先把当前参数转成 protobuf JSON，再发起 gRPC 请求。"
    : "流式接口暂时不支持调用。";

  elements.requestPanel.innerHTML = "";
  const tabs = document.createElement("div");
  tabs.className = "tabs";
  tabs.innerHTML = `
    <button type="button" class="tab ${state.mode === "form" ? "active" : ""}" data-mode="form">表单模式</button>
    <button type="button" class="tab ${state.mode === "json" ? "active" : ""}" data-mode="json">JSON 模式</button>
  `;
  tabs.querySelectorAll(".tab").forEach((button) => {
    button.addEventListener("click", () => {
      state.mode = button.dataset.mode;
      renderRequest();
    });
  });

  elements.requestPanel.appendChild(tabs);

  const layout = document.createElement("div");
  layout.className = "request-layout";

  const left = document.createElement("div");
  const right = document.createElement("div");

  if (state.mode === "form") {
    left.appendChild(renderMessageEditor(detail.requestSchema, []));
  } else {
    left.appendChild(renderJSONEditor());
  }

  const previewTitle = state.mode === "form" ? "当前请求 JSON" : "JSON 模式预览";
  const previewText = state.mode === "form" ? stableStringify(state.formData) : state.requestJson;
  right.innerHTML = `
    <div class="message-card">
      <h3>${escapeHTML(previewTitle)}</h3>
      <p class="muted">这里展示当前会提交出去的请求内容。</p>
      <pre class="code-block">${escapeHTML(previewText)}</pre>
    </div>
  `;

  layout.appendChild(left);
  layout.appendChild(right);
  elements.requestPanel.appendChild(layout);
}

function renderMessageEditor(schema, basePath) {
  const container = document.createElement("div");
  container.className = "message-card";
  container.innerHTML = `
    <h3>${escapeHTML(schema.fullName)}</h3>
    <p class="muted">把需要的字段填上，不填的字段会保持 unset。</p>
  `;

  const fields = document.createElement("div");
  fields.className = "message-fields";

  (schema.fields || []).forEach((field) => {
    fields.appendChild(renderFieldCard(field, basePath));
  });

  if (!(schema.fields || []).length) {
    const empty = document.createElement("p");
    empty.className = "muted";
    empty.textContent = "这个消息没有可编辑字段。";
    fields.appendChild(empty);
  }

  container.appendChild(fields);
  return container;
}

function renderFieldCard(field, basePath) {
  const path = [...basePath, field.jsonName];
  const value = getValueAtPath(state.formData, path);
  const fieldCard = document.createElement("div");
  fieldCard.className = "field-card";

  const badges = [
    `<span class="badge accent">${escapeHTML(field.typeLabel)}</span>`,
    field.repeated ? `<span class="badge">repeated</span>` : "",
    field.map ? `<span class="badge">map</span>` : "",
    field.oneof ? `<span class="badge">oneof: ${escapeHTML(field.oneof)}</span>` : "",
  ]
    .filter(Boolean)
    .join("");

  fieldCard.innerHTML = `
    <div class="field-head">
      <div>
        <div class="field-name">${escapeHTML(field.name)}</div>
        <div class="field-json-name">${escapeHTML(field.jsonName)}</div>
      </div>
      <div class="field-meta">${badges}</div>
    </div>
  `;

  const body = document.createElement("div");
  if (field.map || field.repeated) {
    body.appendChild(renderJSONArrayInput(field, path, value));
  } else if (field.type === "message") {
    body.appendChild(renderNestedMessageInput(field, path, value));
  } else if (field.type === "enum") {
    body.appendChild(renderEnumInput(field, path, value));
  } else if (field.type === "bool") {
    body.appendChild(renderBoolInput(path, value));
  } else if (field.type === "float") {
    body.appendChild(renderScalarInput(path, value, "number", "例如 1.5", { step: "any" }));
  } else if (field.type === "int32") {
    body.appendChild(renderScalarInput(path, value, "number", "例如 42", { step: "1" }));
  } else if (field.type === "int64") {
    body.appendChild(renderScalarInput(path, value, "text", "例如 \"1234567890\""));
    body.appendChild(renderHelper("64 位整数建议按 protobuf JSON 规范使用字符串。"));
  } else if (field.type === "bytes") {
    body.appendChild(renderScalarInput(path, value, "text", "base64 编码"));
  } else {
    body.appendChild(renderScalarInput(path, value, "text", "例如 hello"));
  }

  fieldCard.appendChild(body);
  return fieldCard;
}

function renderScalarInput(path, value, type, placeholder, attrs = {}) {
  const wrapper = document.createElement("label");
  wrapper.className = "field";
  const input = document.createElement("input");
  input.type = type;
  input.placeholder = placeholder;
  input.value = value ?? "";
  Object.entries(attrs).forEach(([key, attrValue]) => input.setAttribute(key, attrValue));

  input.addEventListener("input", () => {
    let nextValue = input.value;
    if (type === "number") {
      if (nextValue === "") {
        deleteAtPath(state.formData, path);
        renderRequest();
        return;
      }
      nextValue = Number(nextValue);
      if (Number.isNaN(nextValue)) {
        return;
      }
    }
    if (nextValue === "") {
      deleteAtPath(state.formData, path);
    } else {
      setValueAtPath(state.formData, path, nextValue);
    }
    refreshRequestPreview();
  });

  wrapper.appendChild(input);
  return wrapper;
}

function renderBoolInput(path, value) {
  const label = document.createElement("label");
  label.className = "toggle";

  const input = document.createElement("input");
  input.type = "checkbox";
  input.checked = value === true;

  input.addEventListener("change", () => {
    if (input.checked) {
      setValueAtPath(state.formData, path, true);
    } else {
      deleteAtPath(state.formData, path);
    }
    refreshRequestPreview();
  });

  const text = document.createElement("span");
  text.textContent = "设为 true";
  label.appendChild(input);
  label.appendChild(text);
  return label;
}

function renderEnumInput(field, path, value) {
  const wrapper = document.createElement("label");
  wrapper.className = "field";
  const select = document.createElement("select");
  select.innerHTML = `<option value="">未设置</option>`;

  (field.enumValues || []).forEach((enumValue) => {
    const option = document.createElement("option");
    option.value = enumValue.name;
    option.textContent = `${enumValue.name} (${enumValue.number})`;
    if (value === enumValue.name) {
      option.selected = true;
    }
    select.appendChild(option);
  });

  select.addEventListener("change", () => {
    if (!select.value) {
      deleteAtPath(state.formData, path);
    } else {
      setValueAtPath(state.formData, path, select.value);
    }
    refreshRequestPreview();
  });

  wrapper.appendChild(select);
  return wrapper;
}

function renderNestedMessageInput(field, path, value) {
  const wrapper = document.createElement("div");
  const actions = document.createElement("div");
  actions.className = "inline-actions";

  const toggleButton = document.createElement("button");
  toggleButton.type = "button";
  toggleButton.textContent = value ? "移除对象" : "启用对象";

  const nested = document.createElement("div");
  nested.className = "nested";
  if (!value) {
    nested.classList.add("hidden");
  }

  toggleButton.addEventListener("click", () => {
    if (getValueAtPath(state.formData, path)) {
      deleteAtPath(state.formData, path);
    } else {
      setValueAtPath(state.formData, path, {});
    }
    renderRequest();
  });

  actions.appendChild(toggleButton);
  wrapper.appendChild(actions);

  if (field.message?.truncated || field.message?.reference) {
    wrapper.appendChild(renderHelper("这个嵌套消息结构较深，建议切到 JSON 模式直接填写。"));
  }

  if (field.message) {
    (field.message.fields || []).forEach((child) => {
      nested.appendChild(renderFieldCard(child, path));
    });
  }

  wrapper.appendChild(nested);
  return wrapper;
}

function renderJSONArrayInput(field, path, value) {
  const wrapper = document.createElement("div");
  const label = document.createElement("label");
  label.className = "field";

  const textarea = document.createElement("textarea");
  textarea.rows = 5;
  textarea.spellcheck = false;
  textarea.placeholder = field.map ? '{\n  "key": "value"\n}' : '[\n  "item1",\n  "item2"\n]';
  textarea.value = value === undefined ? "" : stableStringify(value);

  const helper = renderHelper(
    field.map
      ? `请填写合法 JSON 对象。类型: ${field.mapSpec?.keyTypeLabel} -> ${field.mapSpec?.valueTypeLabel}`
      : "请填写合法 JSON 数组。",
  );

  textarea.addEventListener("input", () => {
    if (!textarea.value.trim()) {
      deleteAtPath(state.formData, path);
      helper.classList.remove("error");
      helper.textContent = field.map
        ? `请填写合法 JSON 对象。类型: ${field.mapSpec?.keyTypeLabel} -> ${field.mapSpec?.valueTypeLabel}`
        : "请填写合法 JSON 数组。";
      refreshRequestPreview();
      return;
    }

    try {
      const parsed = JSON.parse(textarea.value);
      if (field.map && (parsed === null || Array.isArray(parsed) || typeof parsed !== "object")) {
        throw new Error("需要 JSON 对象");
      }
      if (!field.map && !Array.isArray(parsed)) {
        throw new Error("需要 JSON 数组");
      }
      helper.classList.remove("error");
      helper.textContent = field.map
        ? `请填写合法 JSON 对象。类型: ${field.mapSpec?.keyTypeLabel} -> ${field.mapSpec?.valueTypeLabel}`
        : "请填写合法 JSON 数组。";
      setValueAtPath(state.formData, path, parsed);
      refreshRequestPreview();
    } catch (error) {
      helper.classList.add("error");
      helper.textContent = `JSON 解析失败: ${error.message}`;
    }
  });

  label.appendChild(textarea);
  wrapper.appendChild(label);
  wrapper.appendChild(helper);
  return wrapper;
}

function renderJSONEditor() {
  const container = document.createElement("div");
  container.className = "message-card";

  const actions = document.createElement("div");
  actions.className = "inline-actions";

  const syncButton = document.createElement("button");
  syncButton.type = "button";
  syncButton.textContent = "用表单覆盖 JSON";
  syncButton.addEventListener("click", () => {
    state.requestJson = stableStringify(state.formData);
    renderRequest();
  });
  actions.appendChild(syncButton);
  container.appendChild(actions);

  const label = document.createElement("label");
  label.className = "field";
  const textarea = document.createElement("textarea");
  textarea.rows = 18;
  textarea.spellcheck = false;
  textarea.value = state.requestJson;
  textarea.addEventListener("input", () => {
    state.requestJson = textarea.value;
  });
  label.appendChild(textarea);
  container.appendChild(label);
  container.appendChild(renderHelper("JSON 模式适合复杂 repeated / map / 深层嵌套请求。"));
  return container;
}

async function invokeCurrentMethod() {
  if (!state.methodDetail) {
    throw new Error("还没有选择方法");
  }

  let payload;
  if (state.mode === "json") {
    try {
      payload = state.requestJson.trim() ? JSON.parse(state.requestJson) : {};
    } catch (error) {
      throw new Error(`请求 JSON 非法: ${error.message}`);
    }
  } else {
    payload = state.formData;
  }

  try {
    JSON.parse(elements.metadataInput.value || "{}");
  } catch (error) {
    throw new Error(`Metadata JSON 非法: ${error.message}`);
  }

  elements.invokeButton.disabled = true;
  elements.invokeHint.textContent = "正在调用接口...";

  try {
    const response = await fetch("/api/invoke", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        target: elements.targetInput.value.trim(),
        timeoutMs: Number(elements.timeoutInput.value || "10000"),
        authority: elements.authorityInput.value.trim(),
        serverName: elements.serverNameInput.value.trim(),
        useTls: elements.useTlsInput.checked,
        insecureSkipVerify: elements.skipVerifyInput.checked,
        metadata: JSON.parse(elements.metadataInput.value || "{}"),
        fullMethod: state.methodDetail.summary.fullMethod,
        payload,
      }),
    });

    const payloadText = await response.text();
    let result;
    try {
      result = payloadText ? JSON.parse(payloadText) : {};
    } catch {
      result = {
        error: {
          code: `HTTP_${response.status}`,
          message: payloadText || "调用失败",
          raw: payloadText || "调用失败",
        },
      };
    }
    renderResult(result);
  } finally {
    elements.invokeButton.disabled = !state.methodDetail?.unary;
    elements.invokeHint.textContent = state.methodDetail?.unary
      ? "调用前会先把当前参数转成 protobuf JSON，再发起 gRPC 请求。"
      : "流式接口暂时不支持调用。";
  }
}

function renderResult(result) {
  state.result = result;
  if (!result) {
    elements.resultPanel.classList.add("hidden");
    elements.resultPanel.innerHTML = "";
    return;
  }

  elements.resultPanel.classList.remove("hidden");
  const failed = Boolean(result.error);
  const statusText = failed
    ? `${result.error.code || "ERROR"} · ${result.error.message || "调用失败"}`
    : `调用成功 · ${result.durationMs ?? 0} ms`;

  const detailPanels = [];
  if (result.normalizedRequest !== undefined) {
    detailPanels.push(`
      <div class="message-card">
        <h3>Normalized Request</h3>
        <pre class="code-block">${escapeHTML(stableStringify(result.normalizedRequest))}</pre>
      </div>
    `);
  }
  if (result.response !== undefined) {
    detailPanels.push(`
      <div class="message-card">
        <h3>Response</h3>
        <pre class="code-block">${escapeHTML(stableStringify(result.response))}</pre>
      </div>
    `);
  }
  if (failed) {
    detailPanels.push(`
      <div class="message-card">
        <h3>Error</h3>
        <pre class="code-block">${escapeHTML(stableStringify(result.error))}</pre>
      </div>
    `);
  }

  elements.resultPanel.innerHTML = `
    <div class="status">
      <strong class="${failed ? "fail" : "ok"}">${escapeHTML(statusText)}</strong>
      <span class="muted">${escapeHTML(result.fullMethod || state.selectedMethod || "")}</span>
    </div>
    <div class="result-layout">
      ${detailPanels.join("")}
    </div>
  `;
}

function refreshRequestPreview() {
  const preview = elements.requestPanel.querySelector(".code-block");
  if (preview) {
    preview.textContent = stableStringify(state.formData);
  }
}

function renderHelper(text) {
  const helper = document.createElement("div");
  helper.className = "helper";
  helper.textContent = text;
  return helper;
}

function getValueAtPath(root, path) {
  let current = root;
  for (const key of path) {
    if (current == null || typeof current !== "object" || !(key in current)) {
      return undefined;
    }
    current = current[key];
  }
  return current;
}

function setValueAtPath(root, path, value) {
  let current = root;
  for (let index = 0; index < path.length - 1; index += 1) {
    const key = path[index];
    if (current[key] == null || typeof current[key] !== "object" || Array.isArray(current[key])) {
      current[key] = {};
    }
    current = current[key];
  }
  current[path[path.length - 1]] = value;
}

function deleteAtPath(root, path) {
  if (!path.length) {
    return;
  }

  const parents = [];
  let current = root;
  for (let index = 0; index < path.length - 1; index += 1) {
    const key = path[index];
    if (current == null || typeof current !== "object" || !(key in current)) {
      return;
    }
    parents.push([current, key]);
    current = current[key];
  }

  if (current && typeof current === "object") {
    delete current[path[path.length - 1]];
  }

  for (let index = parents.length - 1; index >= 0; index -= 1) {
    const [parent, key] = parents[index];
    const value = parent[key];
    if (value && typeof value === "object" && !Array.isArray(value) && !Object.keys(value).length) {
      delete parent[key];
    }
  }
}

function stableStringify(value) {
  return JSON.stringify(value ?? {}, null, 2);
}

function shortType(fullName) {
  const parts = String(fullName || "").split(".");
  return parts[parts.length - 1] || fullName;
}

function escapeHTML(value) {
  return String(value ?? "")
    .replaceAll("&", "&amp;")
    .replaceAll("<", "&lt;")
    .replaceAll(">", "&gt;")
    .replaceAll('"', "&quot;")
    .replaceAll("'", "&#39;");
}

function renderFatal(message) {
  return `
    <div class="message-card">
      <h2>调试页面没有正常启动</h2>
      <pre class="code-block">${escapeHTML(message)}</pre>
    </div>
  `;
}
