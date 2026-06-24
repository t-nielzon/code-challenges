/*
 * Lazy Chain
 *
 * Create a function `lazyChain` that accepts any value and allows method
 * chaining through the use of `invoke`, delaying execution until the
 * `value` method is called.
 *
 * Example:
 *   let a = lazyChain([1,2,3])
 *               .invoke('map', x => x * 2)
 *               .invoke('reverse');
 *   a.value(); // [6, 4, 2]
 *
 * `invoke` accepts a method name (string) followed by any arguments, and
 * must be general enough to call any prototype method available on the
 * value at the time of execution. Calling `value()` runs every queued
 * method in order and returns the final result.
 */

function lazyChain(value) {
  // your code here
}