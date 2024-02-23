const { render, renderFile } = require('ejs');
const express = require('express');
const app = express();
const port = 3000;
const root = __dirname;

const state = {
    0: '<b class="placeholder"></b>',
    1: '<b class="blueball" hx-boost="true" hx-get="go_down" hx-trigger="every 20ms" ht-target="#balls" hx-swap="innerHTML"></b>',
    2: '<b class="blueball"></b>'
}

let matrix = new Array(10);

app.get('/', (req, res) => {
    matrix = matrix.fill(state[0]);
    matrix[0] = state[1];
    res.sendFile('index.html', {root});
});

app.get('/balls', (req, res) => {
    const balls = matrix.join("");
    res.send(balls);
})

app.get('/go_down', (req, res) => {
    let index = matrix.indexOf(state[1]);
    if (index === matrix.length -1) {
        matrix[index] = state[2];
    } else if (matrix[index + 1] === state[0]) {
        matrix[index] = state[0];
        matrix[index + 1] = state[1];
    } else if (matrix[index + 1] === state[2]) {
        matrix[index] = state[2];
    }
    const balls = matrix.join("");
    res.send(balls);
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port} -> http://localhost:${port}/`);
});

