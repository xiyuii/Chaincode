document.addEventListener("DOMContentLoaded", function() {

    function submitContent() {
        var userInputElement = document.getElementById('userInput');
        var userInput = userInputElement.value;
        var selectedOption = document.getElementById('modelSelect').value;

        // 清空输入框内容
        userInputElement.value = '';

        fetch('/submit', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                user_input: userInput,
                modelSelect: selectedOption
            })
        })
        .then(response => response.json())
        .then(data => {
            var responseDiv = document.getElementById('response');
            responseDiv.innerHTML = data.message;  // 将HTML插入到页面中
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }

    // 监听输入框的 Enter 键按下事件
    var userInputElement = document.getElementById('userInput');
    userInputElement.addEventListener("keydown", function(event) {
        if (event.key === "Enter") {
            event.preventDefault(); // 阻止默认的表单提交行为
            submitContent(); // 调用提交函数
        }
    });

    function goToAnotherPage() {
        window.location.href = '/register';
    }

    function sendDataToBackend() {
        // 定义要发送的数据
        const dataToSend = {
            info: 'Clear'
        };

        // 发送 POST 请求到后端
        fetch('/send_data', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(dataToSend)
        })
        .then(response => response.json())
        .then(data => {
            alert('Data sent to backend successfully: ' + data.message);
            location.reload(); // 刷新页面
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }

    // 将函数绑定到全局，以便在按钮点击时使用
    window.submitContent = submitContent;
    window.goToAnotherPage = goToAnotherPage;
    window.sendDataToBackend = sendDataToBackend;
});
