// Monads: The Maybe Monad
// Difficulty: 5 kyu
//
// Implement the Maybe monad with the following methods:
// - unit: wraps a value in the Just subclass
// - bind: chains Maybe-returning functions together
// - lift: wraps a regular function to return a Maybe, catching exceptions
// - do: chains multiple bound functions in sequence
//
// The Maybe monad has two types: Just (containing a value) and Nothing (empty).

function Maybe() {}

function Just(x) {
  this.toString = function () { return "Just " + x.toString(); };
  this.just = x;
  Object.freeze(this);
}
Just.prototype = new Maybe();
Just.prototype.constructor = Just;

function Nothing() {
  this.toString = function () { return "Nothing"; };
  Object.freeze(this);
}
Nothing.prototype = new Maybe();
Nothing.prototype.constructor = Nothing;

Maybe.unit = function (x) {
  // wraps a value in a Just instance
};

Maybe.bind = function (f) {
  // takes a function that returns a Maybe, returns a function that takes a Maybe
};

Maybe.lift = function (f) {
  // wraps a regular function to return a Maybe, catching exceptions
};

Maybe.do = function (m) {
  // chains multiple bound functions in sequence
};