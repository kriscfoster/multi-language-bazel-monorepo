const express = require('express');
const Calculator = require('../node_calculator/calculator.js');
const app = express();
const calculator = new Calculator();

const port = process.env.PORT | 8080;

app.get('/', (_req, res) => {
  const rand1 = getRandomInt(100);
  const rand2 = getRandomInt(100);
  res.send(`did you know that ${rand1} + ${rand2} = ${calculator.add(rand1, rand2)}`);
});

app.listen(port, () => {
  console.log(`listening on port ${port}`);
});

function getRandomInt(max) {
  return Math.floor(Math.random() * max);
}
