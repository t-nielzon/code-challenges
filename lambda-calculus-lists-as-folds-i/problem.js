/*
 * Lambda Calculus: Lists as folds
 * 
 * In Lambda Calculus, lists can be represented as their right fold.
 * A right fold takes a combining function and an initial value, and returns
 * a single combined value for the whole list.
 * 
 * For example: [ x, y, z ] = c => n => c(x)(c(y)(c(z)(n)))
 * 
 * Booleans are represented as functions that choose between two arguments:
 * True = t => f => t
 * False = t => f => f
 * 
 * The empty list: nil = c => n => n
 * 
 * Task: Implement cons, snoc, map, and filter functions.
 * - cons: adds element to front of list
 * - snoc: adds element to end of list
 * - map: transforms each element in the list
 * - filter: keeps only elements matching a predicate
 */

const cons = element => list => {};

const snoc = element => list => {};

const map = transform => list => {};

const filter = predicate => list => {};