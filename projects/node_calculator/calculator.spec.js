const Calculator = require('./calculator');
const calculator = new Calculator();

it("1 + 2 = 3 >", () => {
  const expected = "Hello World!"
  expect(calculator.add(1, 2)).toEqual(3);
});
