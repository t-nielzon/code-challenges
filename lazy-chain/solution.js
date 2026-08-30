function lazyChain(value) {
  const operations = [];
  
  return {
    invoke(methodName, ...args) {
      operations.push({ methodName, args });
      return this;
    },
    value() {
      let result = value;
      for (const operation of operations) {
        result = result[operation.methodName](...operation.args);
      }
      return result;
    }
  };
}