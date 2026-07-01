(function () {
  // treat the string as an array of chars for delegation to Array.prototype
  function toArray(str) {
    return str.split('');
  }

  // pure (non-mutating) methods return exactly what Array.prototype returns
  String.prototype.map = function (callback, thisArg) {
    return toArray(this).map(callback, thisArg);
  };

  String.prototype.join = function (separator) {
    return toArray(this).join(separator);
  };

  String.prototype.filter = function (callback, thisArg) {
    return toArray(this).filter(callback, thisArg);
  };

  String.prototype.forEach = function (callback, thisArg) {
    return toArray(this).forEach(callback, thisArg);
  };

  String.prototype.some = function (callback, thisArg) {
    return toArray(this).some(callback, thisArg);
  };

  String.prototype.every = function (callback, thisArg) {
    return toArray(this).every(callback, thisArg);
  };

  String.prototype.reduce = function () {
    return Array.prototype.reduce.apply(toArray(this), arguments);
  };

  String.prototype.reduceRight = function () {
    return Array.prototype.reduceRight.apply(toArray(this), arguments);
  };

  // sort and reverse mutate the array but also return it; join back to a string
  String.prototype.sort = function (compareFn) {
    return toArray(this).sort(compareFn).join('');
  };

  String.prototype.reverse = function () {
    return toArray(this).reverse().join('');
  };

  // mutable methods return the resulting string instead of the array's return value
  String.prototype.push = function () {
    var arr = toArray(this);
    arr.push.apply(arr, arguments);
    return arr.join('');
  };

  String.prototype.pop = function () {
    var arr = toArray(this);
    arr.pop();
    return arr.join('');
  };

  String.prototype.shift = function () {
    var arr = toArray(this);
    arr.shift();
    return arr.join('');
  };

  String.prototype.unshift = function () {
    var arr = toArray(this);
    arr.unshift.apply(arr, arguments);
    return arr.join('');
  };

  String.prototype.splice = function () {
    var arr = toArray(this);
    Array.prototype.splice.apply(arr, arguments);
    return arr.join('');
  };
})();