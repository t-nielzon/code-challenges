/*
 * Although shapes can be very different by nature, they can be sorted by the size of their area.
 * 
 * Create different shapes that can be part of a sortable list. The sort order is based on the size of their respective areas:
 * - The area of a Square is the square of its side
 * - The area of a Rectangle is width multiplied by height
 * - The area of a Triangle is base multiplied by height divided by 2
 * - The area of a Circle is the square of its radius multiplied by π
 * - The area of a CustomShape is given
 *
 * The default sort order of a list of shapes is ascending on area size.
 * All classes must be subclasses of the Shape class.
 *
 * Usage example:
 * let side = 1.1234;
 * let radius = 1.1234;
 * let base = 5;
 * let height = 2;
 * let shapes = [ new Square(side), new Circle(radius), new Triangle(base, height) ];
 * shapes.sort( (a,b) => Number(a>b)-Number(a<b) );
 */

class Shape {
  // Base class for all shapes
}

class Square extends Shape {
  constructor(side) {
    // implementation here
  }
}

class Rectangle extends Shape {
  constructor(width, height) {
    // implementation here
  }
}

class Triangle extends Shape {
  constructor(base, height) {
    // implementation here
  }
}

class Circle extends Shape {
  constructor(radius) {
    // implementation here
  }
}

class CustomShape extends Shape {
  constructor(area) {
    // implementation here
  }
}