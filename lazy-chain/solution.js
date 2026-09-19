function lazyChain(value) {
  let operations = [];
  
  return {
    invoke(methodName, ...args) {
      operations.push({ methodName, args });
      return this;
    },
    value() {
      return operations.reduce((val, op) => {
        return val[op.methodName](...op.args);
      }, value);
    }
  };
}