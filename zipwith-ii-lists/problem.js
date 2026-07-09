/*
 * ## Implement zipWith in JavaScript II: Lists :]
 *
 * This zipWith takes a function and two lists and zips the lists together,
 * applying the function to every pair of values.
 * The function value is one new list.
 *
 * If the lists are of unequal length, the output shall be as long as the
 * shorter one. Superfluous values of the longer list are to be ignored.
 * There will be no tests where both list inputs are infinite.
 *
 * Inputs shall not be modified.
 *
 * Lists are built from nodes. A list is represented by its head.
 * Null represents an empty list.
 *
 * function Node(value, next=null) {
 *   this.value = value;
 *   this.next = next;
 * }
 *
 * Examples:
 * zipWith( (a,b) => a+b,
 *          new Node(0, new Node(1, new Node(2))),
 *          new Node(0, new Node(1, new Node(2)))
 * )  =>  new Node(0, new Node(2, new Node(4)))
 *
 * zipWith( Math.pow, list1, list2 )  =>  result_list
 */

function zipWith(fn, list1, list2) {
  
}