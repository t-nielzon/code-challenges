/*
String like [Char] - Difficulty: 5 kyu

In some programming languages, strings internally are implemented like an array of chars.

The objective of this kata is to add to the String.prototype the next Array.prototype methods:

Non-mutating methods:
- Array.prototype.map()
- Array.prototype.join()
- Array.prototype.filter()
- Array.prototype.forEach()
- Array.prototype.some()
- Array.prototype.every()
- Array.prototype.reduce()
- Array.prototype.reduceRight()
- Array.prototype.sort()
- Array.prototype.reverse()

Mutating methods (return modified string instead of array behavior):
- Array.prototype.push()
- Array.prototype.pop()
- Array.prototype.shift()
- Array.prototype.unshift()
- Array.prototype.splice()

Examples:
"Hello".push(" World") // "Hello World"
"Hello".pop() // "Hell"
"Hello".map(c => c.toUpperCase()) // "HELLO"
*/

// Implement String methods that mirror Array prototype methods