function sortOnlyIntegers(arr) {
  const result = JSON.parse(JSON.stringify(arr));
  
  const integers = [];
  
  function extractIntegers(item) {
    if (Array.isArray(item)) {
      item.forEach(el => extractIntegers(el));
    } else if (Number.isInteger(item)) {
      integers.push(item);
    }
  }
  
  extractIntegers(result);
  integers.sort((a, b) => a - b);
  
  let index = 0;
  
  function placeIntegers(item) {
    if (Array.isArray(item)) {
      for (let i = 0; i < item.length; i++) {
        if (Array.isArray(item[i])) {
          placeIntegers(item[i]);
        } else if (Number.isInteger(item[i])) {
          item[i] = integers[index++];
        }
      }
    }
  }
  
  placeIntegers(result);
  return result;
}