function parseMolecule(formula) {
  const stack = [{}];
  let i = 0;
  
  while (i < formula.length) {
    const char = formula[i];
    
    if (char === '(' || char === '[' || char === '{') {
      stack.push({});
      i++;
    } else if (char === ')' || char === ']' || char === '}') {
      const top = stack.pop();
      i++;
      
      let numStr = '';
      while (i < formula.length && /\d/.test(formula[i])) {
        numStr += formula[i];
        i++;
      }
      const multiplier = numStr ? parseInt(numStr) : 1;
      
      for (const element in top) {
        const parent = stack[stack.length - 1];
        parent[element] = (parent[element] || 0) + top[element] * multiplier;
      }
    } else if (/[A-Z]/.test(char)) {
      let element = char;
      i++;
      while (i < formula.length && /[a-z]/.test(formula[i])) {
        element += formula[i];
        i++;
      }
      
      let numStr = '';
      while (i < formula.length && /\d/.test(formula[i])) {
        numStr += formula[i];
        i++;
      }
      const count = numStr ? parseInt(numStr) : 1;
      
      const current = stack[stack.length - 1];
      current[element] = (current[element] || 0) + count;
    } else {
      i++;
    }
  }
  
  return stack[0];
}