function* singleDelta(iterable) {
  let prev = undefined;
  let isFirst = true;
  
  for (const val of iterable) {
    if (!isFirst) {
      yield val - prev;
    }
    prev = val;
    isFirst = false;
  }
}

function* delta(values, level) {
  let current = values;
  
  for (let i = 0; i < level; i++) {
    current = singleDelta(current);
  }
  
  yield* current;
}