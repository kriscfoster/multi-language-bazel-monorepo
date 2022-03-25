const express = require('express');
const Calculator = require('../node_calculator/calculator.js');
const app = express();
const calculator = new Calculator();

const port = process.env.PORT | 8080;

app.get('/', (_req, res) => {
  res.send(`did you know that 1 + 2 = ${calculator.add(1, 2)}`);
});

app.listen(port, () => {
  console.log(`listening on port ${port}`);
});
