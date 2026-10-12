/*
Numbers in JavaScript can be turned into strings. For example, (123.456).toString() 
gives you the string "123.456". But sometimes, numbers are so large or so small that 
the string returned is in a strange format called "scientific notation".

Examples:
- 100 == 1e+2
- 0.01 == 1e-2
- 12.34 == 1.234e+1
- 0.000456 == 4.56e-4

The basic idea is that the coefficient to the left of "e" has its decimal place moved 
a number of places equal to the exponent to the right of "e".

Write a method that can be called on any number that returns the number in its full 
decimal string representation.
*/

Number.prototype.toDecimal = function() {
  
};