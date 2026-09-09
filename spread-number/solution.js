Number.prototype[Symbol.iterator] = function() {
  let n = this.valueOf();
  let current = 1;
  return {
    next: () => {
      if (current <= n) {
        return { value: current++, done: false };
      }
      return { done: true };
    }
  };
};