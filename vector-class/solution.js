function Vector(components) {
  this.components = components.slice();
  this.length = components.length;
}

Vector.prototype._checkLength = function (other) {
  if (this.length !== other.length) {
    throw new Error('Vectors must be of the same length');
  }
};

Vector.prototype.add = function (other) {
  this._checkLength(other);
  return new Vector(this.components.map((v, i) => v + other.components[i]));
};

Vector.prototype.subtract = function (other) {
  this._checkLength(other);
  return new Vector(this.components.map((v, i) => v - other.components[i]));
};

Vector.prototype.dot = function (other) {
  this._checkLength(other);
  return this.components.reduce((sum, v, i) => sum + v * other.components[i], 0);
};

Vector.prototype.norm = function () {
  return Math.sqrt(this.components.reduce((sum, v) => sum + v * v, 0));
};

Vector.prototype.toString = function () {
  return '(' + this.components.join(',') + ')';
};

Vector.prototype.equals = function (other) {
  if (this.length !== other.length) return false;
  return this.components.every((v, i) => v === other.components[i]);
};