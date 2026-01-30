/*
 * Calculating with Functions
 *
 * This time we want to write calculations using functions and get the results.
 *
 * Examples:
 *   seven(times(five()));   // must return 35
 *   four(plus(nine()));     // must return 13
 *   eight(minus(three()));  // must return 5
 *   six(dividedBy(two()));  // must return 3
 *
 * Requirements:
 * - There must be a function for each number from 0 ("zero") to 9 ("nine")
 * - There must be a function for each of the following mathematical operations:
 *   plus, minus, times, dividedBy
 * - Each calculation consist of exactly one operation and two numbers
 * - The most outer function represents the left operand, the most inner
 *   function represents the right operand
 * - Division should be integer division
 */

function zero(op) { return op ? op(0) : 0; }
function one(op) { return op ? op(1) : 1; }
function two(op) { return op ? op(2) : 2; }
function three(op) { return op ? op(3) : 3; }
function four(op) { return op ? op(4) : 4; }
function five(op) { return op ? op(5) : 5; }
function six(op) { return op ? op(6) : 6; }
function seven(op) { return op ? op(7) : 7; }
function eight(op) { return op ? op(8) : 8; }
function nine(op) { return op ? op(9) : 9; }

function plus(n) { return (a) => a + n; }
function minus(n) { return (a) => a - n; }
function times(n) { return (a) => a * n; }
function dividedBy(n) { return (a) => Math.floor(a / n); }