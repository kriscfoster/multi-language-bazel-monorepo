import unittest
from calculator import Calculator

class TestSum(unittest.TestCase):
  def test_sum(self):
    calculator = Calculator()
    self.assertEqual(calculator.add(1, 2), 3)

  def test_sum_negative(self):
    calculator = Calculator()
    self.assertEqual(calculator.add(-4, -11), -15)

if __name__ == '__main__':
  unittest.main()
