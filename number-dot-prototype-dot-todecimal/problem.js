/*
  Number.prototype.toDecimal
  
  Numbers in JavaScript can be turned into strings, but sometimes they use scientific notation.
  For example:
    100 == 1e+2
    0.01 == 1e-2
    12.34 == 1.234e+1
    0.000456 == 4.56e-4
  
  The coefficient (left of 'e') has its decimal place moved by the exponent (right of 'e').
  
  Write a method that can be called on any number that returns the number as a full decimal string,
  without scientific notation.
  
  Examples:
    (1e+2).toDecimal() === "100"
    (1e-2).toDecimal() === "0.01"
    (1.234e+1).toDecimal() === "12.34"
    (4.56e-4).toDecimal() === "0.000456"
*/

Number.prototype.toDecimal = function() {
  
};