/*
 * Parse a linked list from a string
 *
 * Create a function `parse` which accepts exactly one argument `string`
 * which is a string representation of a linked list. Your function must
 * return the corresponding linked list, constructed from instances of
 * the `Node` class.
 *
 * The string representation has the format:
 *   "1 -> 2 -> 3 -> null"
 *
 * Given "1 -> 2 -> 3 -> null", return:
 *   new Node(1, new Node(2, new Node(3)))
 *
 * If the input string is just "null", return null.
 *
 * Node values will always be non-negative integers.
 *
 * Preloaded:
 * class Node {
 *   constructor(data, next = null) {
 *     this.data = data;
 *     this.next = next;
 *   }
 * }
 */

function parse(string) {
}

module.exports = { parse };