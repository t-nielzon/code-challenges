function partitionOn(predicate, items) {
  const falseItems = [];
  const trueItems = [];
  
  for (let item of items) {
    if (predicate(item)) {
      trueItems.push(item);
    } else {
      falseItems.push(item);
    }
  }
  
  const boundaryIndex = falseItems.length;
  items.splice(0, items.length, ...falseItems, ...trueItems);
  
  return boundaryIndex;
}