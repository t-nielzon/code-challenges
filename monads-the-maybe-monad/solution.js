function Maybe() {}

function Just(x) {
  this.toString = function () { return "Just " + x.toString(); };
  this.just = x;
  Object.freeze(this);
}
Just.prototype = new Maybe();
Just.prototype.constructor = Just;

function Nothing() {
  this.toString = function () { return "Nothing"; };
  Object.freeze(this);
}
Nothing.prototype = new Maybe();
Nothing.prototype.constructor = Nothing;

Maybe.unit = function (x) {
  return new Just(x);
};

Maybe.bind = function (f) {
  return function (monadValue) {
    if (!(monadValue instanceof Maybe)) {
      throw new Error("Not a Maybe instance");
    }
    if (monadValue instanceof Just) {
      return f(monadValue.just);
    } else {
      return new Nothing();
    }
  };
};

Maybe.lift = function (f) {
  return function (x) {
    try {
      return new Just(f(x));
    } catch (e) {
      return new Nothing();
    }
  };
};

Maybe.do = function (m) {
  const fns = Array.prototype.slice.call(arguments, 1);
  let result = m;
  for (let i = 0; i < fns.length; i++) {
    result = Maybe.bind(fns[i])(result);
  }
  return result;
};