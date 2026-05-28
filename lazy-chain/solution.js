function lazyChain(value) {
  const operations = [];

  return {
    invoke(methodName, ...args) {
      operations.push({ methodName, args });
      return this;
    },
    value() {
      return operations.reduce(
        (acc, { methodName, args }) => acc[methodName](...args),
        value
      );
    }
  };
}