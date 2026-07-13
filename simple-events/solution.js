function Event() {
  this.handlers = [];
}

Event.prototype.subscribe = function(handler) {
  this.handlers.push(handler);
};

Event.prototype.unsubscribe = function(handler) {
  const index = this.handlers.indexOf(handler);
  if (index !== -1) {
    this.handlers.splice(index, 1);
  }
};

Event.prototype.emit = function(...args) {
  this.handlers.forEach(handler => {
    handler(...args);
  });
};