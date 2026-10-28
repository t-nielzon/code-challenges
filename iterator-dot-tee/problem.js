// Iterator.tee
//
// Define a function that takes an iterator and optionally a number (defaulting to 2),
// and returns an array of $NUMBER forked iterators.
// All returned iterators should iterate the original iterator independently.
// They should also be iterable.
//
// Each forked iterator:
// - Has a next() method that works independently of other forks
// - Can be used with Array.from() or in for...of loops
// - Should cache values from the original iterator so other forks can access them
//
// Example:
// const it = (function*() { yield* [1,2,3,4] })();
// const [xs, ys] = tee(it);
// xs.next() // { value: 1, done: false }
// ys.next() // { value: 1, done: false }
// xs.next() // { value: 2, done: false }

function tee(iterator, n = 2) {
  
}