Function.prototype.wrap = function(wrapper) {
  const original = this;
  return function() {
    return wrapper.apply(this, [original.bind(this), ...arguments]);
  };
};