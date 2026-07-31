function* paperFold() {
  const cache = new Map();
  
  function getValue(n) {
    if (cache.has(n)) return cache.get(n);
    
    let result;
    if (n === 0) {
      result = 1;
    } else {
      const level = Math.ceil(Math.log2(n + 2));
      const sizeOfLevel = (1 << level) - 1;
      const sizeOfPrevLevel = (1 << (level - 1)) - 1;
      const midPos = sizeOfPrevLevel;
      
      if (n === midPos) {
        result = 1;
      } else {
        const mirrorPos = sizeOfLevel - 1 - n;
        result = 1 - getValue(mirrorPos);
      }
    }
    
    cache.set(n, result);
    return result;
  }
  
  let n = 0;
  while (true) {
    yield getValue(n);
    n++;
  }
}