class Vector {
  constructor(components) {
    this.components = components;
  }

  _checkLength(other) {
    if (this.components.length !== other.components.length) {
      throw new Error('Vectors must have the same length');
    }
  }

  add(other) {
    this._checkLength(other);
    return new Vector(this.components.map((v, i) => v + other.components[i]));
  }

  subtract(other) {
    this._checkLength(other);
    return new Vector(this.components.map((v, i) => v - other.components[i]));
  }

  dot(other) {
    this._checkLength(other);
    return this.components.reduce((sum, v, i) => sum + v * other.components[i], 0);
  }

  norm() {
    return Math.sqrt(this.components.reduce((sum, v) => sum + v * v, 0));
  }

  equals(other) {
    if (this.components.length !== other.components.length) return false;
    return this.components.every((v, i) => v === other.components[i]);
  }

  toString() {
    return `(${this.components.join(',')})`;
  }
}