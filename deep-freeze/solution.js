Object.deepFreeze = function (obj) {
  Object.freeze(obj);
  Object.getOwnPropertyNames(obj).forEach(function (prop) {
    const value = obj[prop];
    if (value !== null && typeof value === 'object' && !Object.isFrozen(value)) {
      Object.deepFreeze(value);
    }
  });
  return obj;
};