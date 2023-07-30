function appendToDisplay(text) {
    document.getElementById('display').value += text;
}

function calculate() {
    const expression = document.getElementById('display').value;
    fetch('/calculate', {
        method: 'POST',
        body: expression
    })
    .then(response => response.text())
    .then(result => {
        document.getElementById('display').value = result;
    });
}

function clearDisplay() {
    document.getElementById('display').value = '';
}
