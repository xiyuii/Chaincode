function submitContent() {
    var userInput = document.getElementById('userInput').value;
    var selectedOption = document.getElementById('modelSelect').value;

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
    })
    .catch(error => {
        console.error('Error:', error);
    });
}
