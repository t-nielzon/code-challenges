// A Taste of Curry
// Difficulty: 6 kyu
//
// Your relative has decided to open a small Indian restaurant, and she bought
// an automatic cooker. It publishes an API function:
//
// function cook(dish, nServings, /* other params specific to the dish */) { ... }
//
// The restaurant became popular, and most people order curry with additional
// ingredients. Your relative is tired of typing 'curry' repeatedly, so she asks
// you to write a curry function.
//
// Write a function curry(fun, ...args) that takes another function `fun` and
// an arbitrary number of other arguments, and returns a function that works
// like `fun` with the first arguments replaced by the remaining arguments.
//
// Example:
// function add(a, b, c) { return a+b+c; }
// var addOne = curry(add, 1);
// add(1, 2, 3) === addOne(2, 3); // true
//
// Important: The execution context (this) must be preserved when the returned
// function is invoked.
//
// Example with context:
// var obj = {
//   a: 'foo',
//   b: function (a) { return this.a + a; }
// }
// obj.foobar = curry(obj.b, 'bar');
// obj.foobar() // should return 'foobar'

function curry(fun, ...args) {
    
}