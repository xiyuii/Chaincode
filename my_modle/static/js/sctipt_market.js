document.getElementById('dataForm').addEventListener('submit', function(event) {
    event.preventDefault();
    
    var formData = {
        input1: document.getElementById('input1').value,
        input2: document.getElementById('input2').value,
        input3: document.getElementById('input3').value,
        dropdown: document.getElementById('dropdown').value
    };

    fetch('/submit_api', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
    })
    .then(response => response.json())
    .then(data => {
        alert('Server Response: ' + data.message);
    })
    .catch(error => {
        console.error('Error:', error);
    });
});

document.getElementById('backButton').addEventListener('click', function() {
    window.location.href = '/';
});
