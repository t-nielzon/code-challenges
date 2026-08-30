/*
 * Lazy Chain
 * Difficulty: 5 kyu
 *
 * Method chaining is a useful technique in JavaScript, allowing better
 * composability and readability of functions.
 *
 * One problem with standard JavaScript chaining is that chained methods are
 * immediately executed upon creation, which can open values up to possible
 * mutation before being called.
 *
 * Your task is to create a function lazyChain that accepts an argument and
 * allows chaining, but delays execution to when the value method is called.
 *
 * The function should:
 * - Accept any value
 * - Support method chaining through the invoke method
 * - Queue method calls without executing them
 * - Execute all queued operations when value() is called
 * - Support chaining any prototype methods from standard language
 *
 * Example:
 * let a = lazyChain([1,2,3])
 *   .invoke('map', x => x * 2)
 *   .invoke('reverse');
 *
 * console.log(a.value()); // [6, 4, 2]
 */

function lazyChain(value) {
  
}