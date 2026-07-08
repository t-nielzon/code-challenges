function areEqual(arg1, arg2) {
  // handle null and primitives
  if (arg1 === null || arg2 === null || typeof arg1 !== 'object' || typeof arg2 !== 'object') {
    return arg1 === arg2;
  }
  
  // check if one is array and other is not
  if (Array.isArray(arg1) !== Array.isArray(arg2)) {
    return false;
  }
  
  // handle arrays
  if (Array.isArray(arg1)) {
    if (arg1.length !== arg2.length) return false;
    
    const used = new Array(arg2.length).fill(false);
    
    for (let i = 0; i < arg1.length; i++) {
      let found = false;
      for (let j = 0; j < arg2.length; j++) {
        if (!used[j] && areEqual(arg1[i], arg2[j])) {
          used[j] = true;
          found = true;
          break;
        }
      }
      if (!found) return false;
    }
    return true;
  }
  
  // handle objects
  const keys1 = Object.keys(arg1);
  const keys2 = Object.keys(arg2);
  
  if (keys1.length !== keys2.length) return false;
  
  for (let key of keys1) {
    if (!keys2.includes(key)) return false;
    if (!areEqual(arg1[key], arg2[key])) return false;
  }
  return true;
}