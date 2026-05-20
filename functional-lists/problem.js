/*
 * Functional Lists
 *
 * Create a simple, immutable, singly-linked list with two classes: EmptyList and ListNode.
 *
 * Each class supports: toString(), isEmpty(), length(), push(), remove(), append().
 * If isEmpty() returns false, also: head() and tail().
 *
 * - EmptyList constructor takes no arguments.
 * - ListNode constructor takes (value, next) where next is a ListNode or EmptyList.
 * - toString() => "()" for EmptyList, "(1 2 3)" for a list of 1, 2, 3.
 * - isEmpty() => true for EmptyList, false for ListNode.
 * - length() => number of non-EmptyList nodes.
 * - push(x) => new list with x at the head, sharing as many nodes as possible.
 * - remove(x) => new list with all nodes of value x removed, sharing as many nodes as possible.
 * - append(other) => concatenation of orig and other, sharing as many nodes as possible.
 * - head() => value of first node.
 * - tail() => sublist after first node.
 */

function EmptyList() {
}

EmptyList.prototype.toString = function () {
};

EmptyList.prototype.isEmpty = function () {
};

EmptyList.prototype.length = function () {
};

EmptyList.prototype.push = function (value) {
};

EmptyList.prototype.remove = function (value) {
};

EmptyList.prototype.append = function (other) {
};

function ListNode(value, next) {
}

ListNode.prototype.toString = function () {
};

ListNode.prototype.isEmpty = function () {
};

ListNode.prototype.length = function () {
};

ListNode.prototype.push = function (value) {
};

ListNode.prototype.remove = function (value) {
};

ListNode.prototype.append = function (other) {
};

ListNode.prototype.head = function () {
};

ListNode.prototype.tail = function () {
};