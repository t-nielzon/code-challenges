Object.defineProperty(Array.prototype, 'groupBy', {
  value: function(comparator) {
    const compare = comparator || ((a, b) => a === b);
    const result = [];
    
    for (let i = 0; i < this.length; i++) {
      if (result.length === 0) {
        result.push([this[i]]);
      } else {
        const lastGroup = result[result.length - 1];
        const lastElement = lastGroup[lastGroup.length - 1];
        
        if (compare(this[i], lastElement)) {
          lastGroup.push(this[i]);
        } else {
          result.push([this[i]]);
        }
      }
    }
    
    if (typeof this === 'string') {
      return result.map(group => group.join(''));
    }
    
    return result;
  },
  enumerable: false
});

Object.defineProperty(String.prototype, 'groupBy', {
  value: Array.prototype.groupBy,
  enumerable: false
});