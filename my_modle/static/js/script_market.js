// 获取表单和按钮元素
const form = document.getElementById('dataForm');
const addInputButton = document.getElementById('addInputButton');
const backButton = document.getElementById('backButton');
const headingsAndLabels = document.querySelectorAll('h1, label');

// 创建并插入“使用既有模型”按钮
const resetButton = document.createElement('button');
resetButton.id = 'resetButton';
resetButton.textContent = '使用既有模型';
resetButton.style.display = 'none';  // 初始隐藏
resetButton.style.marginTop = '20px';  // 设置按钮与其他元素的间距
form.appendChild(resetButton);

// 为 "添加新的基础大模型" 按钮添加点击事件监听器
addInputButton.addEventListener('click', function() {
    // 隐藏现有的输入框、下拉框、提交按钮、其他按钮和文本内容
    const inputsAndSelects = form.querySelectorAll('input[type="text"], select');
    inputsAndSelects.forEach(el => el.style.display = 'none');

    // 隐藏标题、标签和按钮
    headingsAndLabels.forEach(el => el.style.display = 'none');
    const submitButton = form.querySelector('input[type="submit"]');
    if (submitButton) submitButton.style.display = 'none';
    addInputButton.style.display = 'none';
    if (backButton) backButton.style.display = 'none';

    // 创建新的输入框（模型名称）
    const newInputName = document.createElement('input');
    newInputName.type = 'text';
    newInputName.name = 'extraModelName';
    newInputName.placeholder = '输入新的模型名称';
    newInputName.required = true;

    // 创建新的文本框（模型调用规范），使用 textarea
    const newInputSpec = document.createElement('textarea');
    newInputSpec.name = 'extraModelSpec';
    newInputSpec.placeholder = '输入模型调用规范';
    newInputSpec.required = true;
    newInputSpec.style.width = '100%';  // 占满父容器宽度
    newInputSpec.style.minHeight = '50px';  // 初始高度
    newInputSpec.style.resize = 'vertical';  // 允许用户手动调整高度

    // 创建一个包含新输入框的容器
    const newInputContainer = document.createElement('div');
    newInputContainer.style.marginBottom = '20px';  // 给新输入框添加底部间距
    newInputContainer.appendChild(newInputName);
    newInputContainer.appendChild(newInputSpec);

    // 将新输入框的容器添加到表单
    form.appendChild(newInputContainer);

    // 显示新的提交按钮
    const newSubmitButton = document.createElement('input');
    newSubmitButton.type = 'submit';
    newSubmitButton.value = '提交新模型';
    newSubmitButton.style.marginTop = '20px';
    form.appendChild(newSubmitButton);

    // 显示“使用既有模型”按钮
    resetButton.style.display = 'block';
});

// 处理表单提交事件
form.addEventListener('submit', function(event) {
    event.preventDefault();  // 阻止表单的默认提交行为

    // 获取表单数据
    const formData = {
        input1: document.getElementById('input1') ? document.getElementById('input1').value : '',
        input2: document.getElementById('input2') ? document.getElementById('input2').value : '',
        input3: document.getElementById('input3') ? document.getElementById('input3').value : '',
        input4: document.getElementById('input4') ? document.getElementById('input4').value : '',
        dropdown: document.getElementById('dropdown') ? document.getElementById('dropdown').value : ''
    };

    // 如果额外的输入框存在，则加入到表单数据
    const extraModelName = document.querySelector('input[name="extraModelName"]');
    const extraModelSpec = document.querySelector('textarea[name="extraModelSpec"]');

    if (extraModelName && extraModelSpec) {
        formData.extraModelName = extraModelName.value;
        formData.extraModelSpec = extraModelSpec.value;
    }

    // 使用 fetch 发送 POST 请求
    fetch('/submit_api', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
    })
    .then(response => response.json())
    .then(data => {
        alert('服务器响应: ' + data.message);  // 显示服务器响应
    })
    .catch(error => {
        console.error('Error:', error);  // 处理错误
    });
});

// 处理 "返回" 按钮点击事件
backButton.addEventListener('click', function() {
    window.location.href = '/';  // 重定向到首页
});

// 处理“使用既有模型”按钮点击事件
resetButton.addEventListener('click', function() {
    // 显示现有的输入框、下拉框、提交按钮、其他按钮和文本内容
    const inputsAndSelects = form.querySelectorAll('input[type="text"], select');
    inputsAndSelects.forEach(el => el.style.display = 'block');
    headingsAndLabels.forEach(el => el.style.display = 'block');
    const originalSubmitButton = form.querySelector('input[type="submit"]');
    if (originalSubmitButton) originalSubmitButton.style.display = 'block';
    addInputButton.style.display = 'block';
    if (backButton) backButton.style.display = 'block';

    // 隐藏新输入框和提交按钮
    const newInputs = form.querySelectorAll('input[name="extraModelName"], textarea[name="extraModelSpec"]');
    newInputs.forEach(el => el.parentNode.remove());  // 删除新输入框容器
    const newSubmitButton = form.querySelector('input[type="submit"][value="提交新模型"]');
    if (newSubmitButton) newSubmitButton.remove();

    // 隐藏“使用既有模型”按钮
    resetButton.style.display = 'none';
});
