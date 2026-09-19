/**
 * Lazy Chain
 * Difficulty: 5 kyu
 * 
 * Method chaining is a useful technique in JavaScript, allowing better composability
 * and readability of functions. However, standard JavaScript chaining executes methods
 * immediately upon creation.
 * 
 * To combat this, create a function `lazyChain` that accepts an argument and allows
 * chaining through the `invoke` method, but delays execution until the `value` method
 * is called.
 * 
 * Example:
 * let a = lazyChain([1,2,3])
 *                     .invoke('map', x => x * 2)
 *                     .invoke('reverse');
 * 
 * console.log(a.value()); // [6, 4, 2]
 * 
 * Your task is to create the `lazyChain` function which:
 * - Accepts any initial value
 * - Allows method chaining through the `invoke` method
 * - Delays execution until `value()` is called
 * - Supports any prototype methods from the standard language
 */

function lazyChain(value) {
  // Your code here
}