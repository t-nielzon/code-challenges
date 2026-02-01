function Lazy() {
  this.chain = [];
}

Lazy.prototype.add = function(fn /*, arg1, arg2, ... */) {
  var args = Array.prototype.slice.call(arguments, 1);
  this.chain.push({ fn: fn, args: args });
  return this;
};

Lazy.prototype.invoke = function(target) {
  var result = target;
  
  for (var i = 0; i < this.chain.length; i++) {
    var step = this.chain[i];
    var args = step.args.concat(result);
    result = step.fn.apply(null, args);
    
    if (!Array.isArray(result)) {
      if (i < this.chain.length - 1) {
        result = [result];
      }
    }
  }
  
  return result;
};