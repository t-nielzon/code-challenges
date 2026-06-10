Array.prototype.foldr = function (fn, z) {
  const self = this;
  const go = (i) => {
    if (i >= self.length) return z;
    // the rest of the fold is wrapped in a thunk so it is only computed
    // when fn actually asks for it; memoized so repeated access is cheap
    let cached, done = false;
    const rest = () => {
      if (!done) { cached = go(i + 1); done = true; }
      return cached;
    };
    return fn(self[i], rest);
  };
  return go(0);
};

String.prototype.foldr = Array.prototype.foldr;