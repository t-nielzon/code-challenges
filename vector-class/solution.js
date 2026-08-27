// solution.js
function Vector(components) {
  this.components = components;
}

Vector.prototype.add = function(vector) {
  if (this.components.length !== vector.components.length) {
    throw new Error('Vector dimensions must match');
  }
  const result = this.components.map((val, i) => val + vector.components[i]);
  return new Vector(result);
};

Vector.prototype.subtract = function(vector) {
  if (this.components.length !== vector.components.length) {
    throw new Error('Vector dimensions must match');
  }
  const result = this.components.map((val, i) => val - vector.components[i]);
  return new Vector(result);
};

Vector.prototype.dot = function(vector) {
  if (this.components.length !== vector.components.length) {
    throw new Error('Vector dimensions must match');
  }
  return this.components.reduce((sum, val, i) => sum + val * vector.components[i], 0);
};

Vector.prototype.norm = function() {
  return Math.sqrt(this.components.reduce((sum, val) => sum + val * val, 0));
};

Vector.prototype.toString = function() {
  return '(' + this.components.join(',') + ')';
};

Vector.prototype.equals = function(vector) {
  if (this.components.length !== vector.components.length) {
    return false;
  }
  return this.components.every((val, i) => val === vector.components[i]);
};