document.addEventListener("DOMContentLoaded", function() {

    function submitContent() {
        var userInputElement = document.getElementById('userInput');
        var userInput = userInputElement.value;
        var selectedOption = document.getElementById('modelSelect').value;

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
            responseDiv.innerHTML = data.message;
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }

    var userInputElement = document.getElementById('userInput');
    userInputElement.addEventListener("keydown", function(event) {
        if (event.key === "Enter") {
            event.preventDefault();
            submitContent();
        }
    });

    function goToAnotherPage() {
        window.location.href = '/register';
    }

    function sendDataToBackend() {
        const dataToSend = {
            info: 'Clear'
        };

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
            location.reload();
        })
        .catch(error => {
            console.error('Error:', error);
        });
    }

    window.submitContent = submitContent;
    window.goToAnotherPage = goToAnotherPage;
    window.sendDataToBackend = sendDataToBackend;
});
