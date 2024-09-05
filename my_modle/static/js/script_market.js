const form = document.getElementById('dataForm');
const addInputButton = document.getElementById('addInputButton');
const backButton = document.getElementById('backButton');
const headingsAndLabels = document.querySelectorAll('h1, label');

const resetButton = document.createElement('button');
resetButton.id = 'resetButton';
resetButton.textContent = '使用既有模型';
resetButton.style.display = 'none';  // 初始隐藏
resetButton.style.marginTop = '20px';
form.appendChild(resetButton);

addInputButton.addEventListener('click', function() {
    const inputsAndSelects = form.querySelectorAll('input[type="text"], select');
    inputsAndSelects.forEach(el => el.style.display = 'none');

    headingsAndLabels.forEach(el => el.style.display = 'none');
    const submitButton = form.querySelector('input[type="submit"]');
    if (submitButton) submitButton.style.display = 'none';
    addInputButton.style.display = 'none';
    if (backButton) backButton.style.display = 'none';

    const newInputName = document.createElement('input');
    newInputName.type = 'text';
    newInputName.name = 'extraModelName';
    newInputName.placeholder = '输入新的模型名称';
    newInputName.required = true;

    const newInputSpec = document.createElement('textarea');
    newInputSpec.name = 'extraModelSpec';
    newInputSpec.placeholder = '输入模型调用规范';
    newInputSpec.required = true;
    newInputSpec.style.width = '100%'; 
    newInputSpec.style.minHeight = '50px';
    newInputSpec.style.resize = 'vertical';

    const newInputContainer = document.createElement('div');
    newInputContainer.style.marginBottom = '20px';
    newInputContainer.appendChild(newInputName);
    newInputContainer.appendChild(newInputSpec);

    form.appendChild(newInputContainer);

    const newSubmitButton = document.createElement('input');
    newSubmitButton.type = 'submit';
    newSubmitButton.value = '提交新模型';
    newSubmitButton.style.marginTop = '20px';
    form.appendChild(newSubmitButton);

    resetButton.style.display = 'block';
});

form.addEventListener('submit', function(event) {
    event.preventDefault();

    const formData = {
        input1: document.getElementById('input1') ? document.getElementById('input1').value : '',
        input2: document.getElementById('input2') ? document.getElementById('input2').value : '',
        input3: document.getElementById('input3') ? document.getElementById('input3').value : '',
        input4: document.getElementById('input4') ? document.getElementById('input4').value : '',
        dropdown: document.getElementById('dropdown') ? document.getElementById('dropdown').value : ''
    };

    const extraModelName = document.querySelector('input[name="extraModelName"]');
    const extraModelSpec = document.querySelector('textarea[name="extraModelSpec"]');

    if (extraModelName && extraModelSpec) {
        formData.extraModelName = extraModelName.value;
        formData.extraModelSpec = extraModelSpec.value;
    }

    fetch('/submit_api', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
    })
    .then(response => response.json())
    .then(data => {
        alert('服务器响应: ' + data.message); 
    })
    .catch(error => {
        console.error('Error:', error); 
    });
});

backButton.addEventListener('click', function() {
    window.location.href = '/';  // 重定向到首页
});

resetButton.addEventListener('click', function() {
    const inputsAndSelects = form.querySelectorAll('input[type="text"], select');
    inputsAndSelects.forEach(el => el.style.display = 'block');
    headingsAndLabels.forEach(el => el.style.display = 'block');
    const originalSubmitButton = form.querySelector('input[type="submit"]');
    if (originalSubmitButton) originalSubmitButton.style.display = 'block';
    addInputButton.style.display = 'block';
    if (backButton) backButton.style.display = 'block';

    const newInputs = form.querySelectorAll('input[name="extraModelName"], textarea[name="extraModelSpec"]');
    newInputs.forEach(el => el.parentNode.remove());
    const newSubmitButton = form.querySelector('input[type="submit"][value="提交新模型"]');
    if (newSubmitButton) newSubmitButton.remove();

    resetButton.style.display = 'none';
});
