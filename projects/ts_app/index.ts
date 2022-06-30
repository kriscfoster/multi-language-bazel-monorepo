import Calculator from './calculator';

const calculator: Calculator = new Calculator();
const num1: number = 1;
const num2: number = 2;

const result = calculator.add(num1, num2);

console.log(`${num1} + ${num2} = ${result}`);
