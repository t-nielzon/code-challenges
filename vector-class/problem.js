// problem.js
/*
Create a Vector object that supports addition, subtraction, dot products, and norms.

Examples:
a = new Vector([1, 2, 3])
b = new Vector([3, 4, 5])
c = new Vector([5, 6, 7, 8])

a.add(b)      // should return a new Vector([4, 6, 8])
a.subtract(b) // should return a new Vector([-2, -2, -2])
a.dot(b)      // should return 1*3 + 2*4 + 3*5 = 26
a.norm()      // should return sqrt(1^2 + 2^2 + 3^2) = sqrt(14)
a.add(c)      // throws an error

Methods required:
- add(vector): returns a new Vector with element-wise addition
- subtract(vector): returns a new Vector with element-wise subtraction
- dot(vector): returns the dot product (scalar)
- norm(): returns the magnitude of the vector
- toString(): returns string like "(1,2,3)"
- equals(vector): checks if two vectors are equal

Must throw an error if operations are performed on vectors of different lengths.
*/

function Vector(components) {
}