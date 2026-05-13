[
  'map', 'filter', 'forEach', 'some', 'every',
  'reduce', 'reduceRight', 'sort', 'reverse', 'join'
].forEach(function (method) {
  String.prototype[method] = function () {
    return Array.prototype[method].apply(this.split(''), arguments);
  };
});

String.prototype.push = function () {
  return this + Array.prototype.slice.call(arguments).join('');
};

String.prototype.pop = function () {
  return this.slice(0, -1);
};

String.prototype.shift = function () {
  return this.slice(1);
};

String.prototype.unshift = function () {
  return Array.prototype.slice.call(arguments).join('') + this;
};

String.prototype.splice = function (start, deleteCount) {
  var arr = this.split('');
  var args = [start, deleteCount].concat(Array.prototype.slice.call(arguments, 2));
  Array.prototype.splice.apply(arr, args);
  return arr.join('');
};