async function submitContent() {
    const userInput = document.getElementById('userInput').value;

    try {
        const response = await fetch('/submit', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ user_input: userInput })
        });
        
        const data = await response.json();
        document.getElementById('response').innerHTML = `
            <p>Apply: ${data.message}</p>
        `;
    } catch (error) {
        console.error('Error:', error);
    }
}
