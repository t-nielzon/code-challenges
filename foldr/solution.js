Array.prototype.foldr = function(fn, z) {
  if (this.length === 0) {
    return z;
  }
  return fn(this[0], () => this.slice(1).foldr(fn, z));
};

String.prototype.foldr = function(fn, z) {
  if (this.length === 0) {
    return z;
  }
  return fn(this[0], () => this.slice(1).foldr(fn, z));
};