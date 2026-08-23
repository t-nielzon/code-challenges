/*
Calculating with Objects

Bob made a class called Num to work with arithmetic operations.
He wants to perform calculations like:
  x + y == 105
  x * y == 500
  x - y == 95
  x / y == 20
  Math.floor(new Num(100.5)) == 100

But it doesn't work - he keeps getting '[object Object][object Object]' as a result.

We need to fix the class so arithmetic operations work while keeping toString() 
and valueOf() methods functional for their original purposes.
*/

class Num {
  constructor(num){
    this.num=num;
  }
  toString() {
    return "The number is " + this.num;
  }
  valueOf() {
    return {num: this.num};
  }
}