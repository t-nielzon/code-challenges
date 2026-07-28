// Calculating with objects
// 
// Bob made a class called Num and is trying to use it with arithmetic operations.
// Currently, when he tries to do operations like x + y, he gets '[object Object][object Object]'
// instead of the expected numeric result.
//
// The issue is that valueOf() returns an object instead of a primitive value,
// which prevents JavaScript's type coercion from working correctly with arithmetic operations.
//
// Task: Fix the Num class so that arithmetic operations work (x + y == 105, x * y == 500, etc.)
// while keeping the toString() and valueOf() methods functional.

class Num {
  constructor(num){
    this.num = num;
  }
  
  toString() {
    // TODO: implement
  }
  
  valueOf() {
    // TODO: implement
  }
}