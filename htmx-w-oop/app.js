const express = require('express');
const path = require('path');
const app = express();

const Button = require('./objects');

const root = __dirname;

const PORT = process.env.PORT || 3000;

app.use(express.static('public'));

app.get('/', (req, res) => {
    res.sendFile('index.html', {root});
  });

app.get('/blue-button', (req, res) => {
    let props = {
        "text": "Clique aqui!",
        "class": "p-1 bg-blue-600",
        "id": "blue-button",
        "hx-target": "this",
        "hx-get": "/red-button",
        "hx-trigger": "click",
        "hx-swap": "outerHTML"
    };
    res.send(new Button(props).html);
});

app.get('/red-button',  (req, res) => {
    let props = {
        "text": "Clique aqui!",
        "class": "p-2 bg-red-600 rounded-lg",
        "id": "blue-button",
        "hx-target": "this",
        "hx-get": "/blue-button",
        "hx-trigger": "click",
        "hx-swap": "outerHTML"
    };
    res.send(new Button(props).html);
});

app.listen(PORT, () => {
  console.log(`Servidor rodando em http://localhost:${PORT}`);
});