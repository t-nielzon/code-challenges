Array.prototype.permutations = function(n) {
  if (!Number.isInteger(n) || n < 0) {
    throw new Error('Invalid input');
  }
  
  if (n > this.length) {
    return [];
  }
  
  if (n === 0) {
    return [[]];
  }
  
  const result = [];
  
  function permute(arr, remaining, current) {
    if (remaining === 0) {
      result.push([...current]);
      return;
    }
    
    for (let i = 0; i < arr.length; i++) {
      const item = arr[i];
      const newArr = arr.slice(0, i).concat(arr.slice(i + 1));
      permute(newArr, remaining - 1, [...current, item]);
    }
  }
  
  permute(this, n, []);
  return result;
};