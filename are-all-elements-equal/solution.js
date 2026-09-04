function eqAll() {
  const iterator = this[Symbol.iterator]();
  const firstResult = iterator.next();
  
  if (firstResult.done) return true;
  
  const firstValue = firstResult.value;
  let result;
  while (!(result = iterator.next()).done) {
    if (result.value !== firstValue) return false;
  }
  return true;
}

Object.defineProperty(String.prototype, 'eqAll', { value: eqAll, enumerable: false });
Object.defineProperty(Array.prototype, 'eqAll', { value: eqAll, enumerable: false });
Object.defineProperty(Set.prototype, 'eqAll', { value: eqAll, enumerable: false });