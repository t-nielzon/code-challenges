String.prototype.map = function(callback, thisArg) {
  return Array.from(this).map(callback, thisArg).join('');
};

String.prototype.filter = function(callback, thisArg) {
  return Array.from(this).filter(callback, thisArg).join('');
};

String.prototype.join = function(separator = ',') {
  return Array.from(this).join(separator);
};

String.prototype.forEach = function(callback, thisArg) {
  Array.from(this).forEach(callback, thisArg);
};

String.prototype.some = function(callback, thisArg) {
  return Array.from(this).some(callback, thisArg);
};

String.prototype.every = function(callback, thisArg) {
  return Array.from(this).every(callback, thisArg);
};

String.prototype.reduce = function(callback, initialValue) {
  return Array.from(this).reduce(callback, initialValue);
};

String.prototype.reduceRight = function(callback, initialValue) {
  return Array.from(this).reduceRight(callback, initialValue);
};

String.prototype.sort = function(compareFn) {
  return Array.from(this).sort(compareFn).join('');
};

String.prototype.reverse = function() {
  return Array.from(this).reverse().join('');
};

String.prototype.push = function(...elements) {
  return this + elements.join('');
};

String.prototype.pop = function() {
  return this.slice(0, -1);
};

String.prototype.shift = function() {
  return this.slice(1);
};

String.prototype.unshift = function(...elements) {
  return elements.join('') + this;
};

String.prototype.splice = function(start, deleteCount, ...items) {
  const arr = Array.from(this);
  arr.splice(start, deleteCount, ...items);
  return arr.join('');
};