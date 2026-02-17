var findInArray = function(array, fn) {
  for (var i = 0; i < array.length; i++) {
    if (fn(array[i], i)) return i;
  }
  return -1;
};