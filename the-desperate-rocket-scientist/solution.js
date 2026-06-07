function fix_countdown() {
  // jeff polluted Array.prototype with enumerable properties, which a for...in
  // loop walks alongside the real indices. making every inherited prototype
  // property non-enumerable restores the original countdown behavior.
  Object.getOwnPropertyNames(Array.prototype).forEach(function (key) {
    Object.defineProperty(Array.prototype, key, { enumerable: false });
  });
}