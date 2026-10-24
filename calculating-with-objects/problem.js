// Calculating with Objects
// Difficulty: 5 kyu
//
// Bob created a Num class with toString() and valueOf() methods.
// He wants to use it with arithmetic operations:
//   const x = new Num(100);
//   const y = new Num(5);
//   x + y == 105;
//   x*y == 500;
//   x-y == 95
//   x/y == 20
//   Math.floor(new Num(100.5)) == 100
//
// But arithmetic returns '[object Object][object Object]' instead.
// The valueOf() and toString() methods must remain functional,
// but the class needs to support arithmetic operations.

class Num {
  constructor(num) {
    this.num = num;
  }
  
  toString() {
    return "The number is " + this.num;
  }
  
  valueOf() {
    // Fix needed here to make arithmetic work
  }
}