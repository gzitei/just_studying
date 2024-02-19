const express = require('express');
const path = require('path');
const app = express();

const PORT = process.env.PORT || 3000;

app.get('/', (req, res) => {
  res.sendFile(path.join(__dirname, 'index.html'));
});

let count = 0;

app.get('/increment', (req, res) => {
    count++;
    console.log("count incremented", count);
    res.status(200).send(count.toString());
});

app.get('/decrement', (req, res) => {
    count--;
    console.log("count decremented:", count);
    res.status(200).send(count.toString());
});

app.get('/setCounter', (req, res) => {
    console.log("count reseted!");
    res.status(200).send(count.toString());
});

app.listen(PORT, () => {
  console.log(`Servidor rodando em http://localhost:${PORT}`);
});