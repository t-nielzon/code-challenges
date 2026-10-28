function tee(iterator, n = 2) {
  const cache = [];
  let exhausted = false;
  let doneResult = null;
  
  const pullNext = () => {
    if (exhausted) {
      return doneResult;
    }
    const result = iterator.next();
    if (result.done) {
      exhausted = true;
      doneResult = result;
      return result;
    }
    cache.push(result.value);
    return result;
  };
  
  const forks = [];
  for (let i = 0; i < n; i++) {
    const fork = {
      index: 0,
      next() {
        if (this.index < cache.length) {
          return { value: cache[this.index++], done: false };
        }
        const result = pullNext();
        if (!result.done) {
          this.index++;
        }
        return result;
      },
      [Symbol.iterator]() {
        return this;
      }
    };
    forks.push(fork);
  }
  
  return forks;
}