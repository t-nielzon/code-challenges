var Vector = function(components) {
  this.components = components;
};

Vector.prototype.add = function(other) {
  if (this.components.length !== other.components.length) throw new Error('Vectors must have the same length');
  return new Vector(this.components.map((v, i) => v + other.components[i]));
};

Vector.prototype.subtract = function(other) {
  if (this.components.length !== other.components.length) throw new Error('Vectors must have the same length');
  return new Vector(this.components.map((v, i) => v - other.components[i]));
};

Vector.prototype.dot = function(other) {
  if (this.components.length !== other.components.length) throw new Error('Vectors must have the same length');
  return this.components.reduce((sum, v, i) => sum + v * other.components[i], 0);
};

Vector.prototype.norm = function() {
  return Math.sqrt(this.components.reduce((sum, v) => sum + v * v, 0));
};

Vector.prototype.toString = function() {
  return '(' + this.components.join(',') + ')';
};

Vector.prototype.equals = function(other) {
  return this.components.length === other.components.length &&
    this.components.every((v, i) => v === other.components[i]);
};