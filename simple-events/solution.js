function Event() {
  this.handlers = [];
}

Event.prototype.subscribe = function(fn) {
  this.handlers.push(fn);
};

Event.prototype.unsubscribe = function(fn) {
  this.handlers = this.handlers.filter(handler => handler !== fn);
};

Event.prototype.emit = function(...args) {
  this.handlers.forEach(handler => handler(...args));
};