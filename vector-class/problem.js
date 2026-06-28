/*
 * Vector class
 * Difficulty: 5 kyu
 *
 * Create a Vector object that supports addition, subtraction, dot products, and norms.
 *
 * var a = new Vector([1, 2, 3]);
 * var b = new Vector([3, 4, 5]);
 * var c = new Vector([5, 6, 7, 8]);
 *
 * a.add(b);      // should return a new Vector([4, 6, 8])
 * a.subtract(b); // should return a new Vector([-2, -2, -2])
 * a.dot(b);      // should return 1*3 + 2*4 + 3*5 = 26
 * a.norm();      // should return sqrt(1^2 + 2^2 + 3^2) = sqrt(14)
 * a.add(c);      // throws an error
 *
 * If you try to add, subtract, or dot two vectors with different lengths, you must throw an error!
 *
 * Also provide:
 * - toString: a.toString() === '(1,2,3)'
 * - equals: to check that two vectors that have the same components are equal
 */

class Vector {
  constructor(components) {}
  add(other) {}
  subtract(other) {}
  dot(other) {}
  norm() {}
  equals(other) {}
  toString() {}
}