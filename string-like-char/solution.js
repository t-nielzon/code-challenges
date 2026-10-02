String.prototype.map = function(callback) {
  return Array.prototype.map.call(this, callback).join('');
};

String.prototype.join = function(separator) {
  return Array.prototype.join.call(this, separator);
};

String.prototype.filter = function(callback) {
  return Array.prototype.filter.call(this, callback).join('');
};

String.prototype.forEach = function(callback) {
  Array.prototype.forEach.call(this, callback);
};

String.prototype.some = function(callback) {
  return Array.prototype.some.call(this, callback);
};

String.prototype.every = function(callback) {
  return Array.prototype.every.call(this, callback);
};

String.prototype.reduce = function(callback, initialValue) {
  if (arguments.length > 1) {
    return Array.prototype.reduce.call(this, callback, initialValue);
  }
  return Array.prototype.reduce.call(this, callback);
};

String.prototype.reduceRight = function(callback, initialValue) {
  if (arguments.length > 1) {
    return Array.prototype.reduceRight.call(this, callback, initialValue);
  }
  return Array.prototype.reduceRight.call(this, callback);
};

String.prototype.sort = function(callback) {
  return Array.prototype.sort.call(this, callback).join('');
};

String.prototype.reverse = function() {
  return Array.prototype.reverse.call(this).join('');
};

String.prototype.push = function(...args) {
  return this + args.join('');
};

String.prototype.pop = function() {
  return this.slice(0, -1);
};

String.prototype.shift = function() {
  return this.slice(1);
};

String.prototype.unshift = function(...args) {
  return args.join('') + this;
};

String.prototype.splice = function(start, deleteCount, ...items) {
  const arr = this.split('');
  arr.splice(start, deleteCount, ...items);
  return arr.join('');
};