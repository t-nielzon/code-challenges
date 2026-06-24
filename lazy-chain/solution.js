function lazyChain(value) {
  const queue = [];

  return {
    invoke(methodName, ...args) {
      queue.push({ methodName, args });
      return this;
    },
    value() {
      return queue.reduce(
        (acc, { methodName, args }) => acc[methodName](...args),
        value
      );
    }
  };
}