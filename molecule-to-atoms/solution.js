function parseMolecule(formula) {
  const stack = [{}];
  let i = 0;

  while (i < formula.length) {
    const char = formula[i];

    if (char === '(' || char === '[' || char === '{') {
      stack.push({});
      i++;
    } else if (char === ')' || char === ']' || char === '}') {
      i++;
      let num = '';
      while (i < formula.length && /\d/.test(formula[i])) {
        num += formula[i];
        i++;
      }
      const multiplier = num ? parseInt(num) : 1;

      const popped = stack.pop();
      const current = stack[stack.length - 1];

      for (const [element, count] of Object.entries(popped)) {
        current[element] = (current[element] || 0) + count * multiplier;
      }
    } else if (/[A-Z]/.test(char)) {
      let element = char;
      i++;
      while (i < formula.length && /[a-z]/.test(formula[i])) {
        element += formula[i];
        i++;
      }

      let num = '';
      while (i < formula.length && /\d/.test(formula[i])) {
        num += formula[i];
        i++;
      }
      const count = num ? parseInt(num) : 1;

      const current = stack[stack.length - 1];
      current[element] = (current[element] || 0) + count;
    } else {
      i++;
    }
  }

  return stack[0];
}