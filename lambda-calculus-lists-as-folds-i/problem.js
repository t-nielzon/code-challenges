/*
Lambda Calculus: Lists as folds

In Lambda Calculus, lists can be represented as their right fold. A right fold
takes a combining function and an initial value, and returns a single combined
value for the whole list.

In JavaScript syntax:
[ x, y, z ] = c => n => c(x)(c(y)(c(z)(n))) ;

Booleans are represented as functions that choose between two arguments.
This encoding uses Church Booleans and Church-encoded lists.

Task:
Write functions cons and snoc. Both take an element and a list and return a new
list with the element added to the list, in first and last position respectively.

Write function map. It takes a transforming function and a list and returns a
new list with every element transformed.

Write function filter. It takes a predicate and a list and returns a new,
filtered list. A predicate is a function that takes a value and returns a Boolean.

Preloaded:
False = t => f => f                             // Church Boolean
True  = t => f => t                             // Church Boolean

nil = c => n => n                               // constant: the empty list
isEmpty = xs => xs ( _ => _ => False ) (True)   // returns a Church Boolean indicating if a list is empty
head    = xs => xs ( x => _ => x ) (undefined)  // returns the first element of a list
tail    = xs => something quite clever          // returns a list without its first element
*/

const cons = undefined;
const snoc = undefined;
const map = undefined;
const filter = undefined;