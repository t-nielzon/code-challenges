function fix_countdown() {
  // jeff polluted Array.prototype / Object.prototype with enumerable
  // properties, which leak into the `for...in` loop inside countdown().
  // making every inherited property non-enumerable restores the loop.
  [Array.prototype, Object.prototype].forEach(function (proto) {
    Object.getOwnPropertyNames(proto).forEach(function (prop) {
      try {
        Object.defineProperty(proto, prop, { enumerable: false });
      } catch (e) {
        // some built-in props are non-configurable; ignore them.
      }
    });
  });
}