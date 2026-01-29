function listComprehension(str) {
  // Trim and validate basic structure
  str = str.trim();
  
  if (!str.startsWith('[') || !str.endsWith(']')) {
    throw new Error('Invalid list comprehension');
  }
  
  // Remove outer brackets
  const inner = str.slice(1, -1).trim();
  
  // Parse the list comprehension: expression for variable in iterable
  const forMatch = inner.match(/^(.+?)\s+for\s+(\w+)\s+in\s+(.+)$/);
  
  if (!forMatch) {
    throw new Error('Invalid list comprehension syntax');
  }
  
  const [, expression, variable, iterableExpr] = forMatch;
  
  // Evaluate the iterable
  let iterable;
  try {
    iterable = eval(iterableExpr);
  } catch (e) {
    throw new Error('Invalid iterable expression');
  }
  
  if (!Array.isArray(iterable)) {
    throw new Error('Iterable must be an array');
  }
  
  // Map over the iterable, evaluating the expression for each element
  const result = [];
  for (const value of iterable) {
    try {
      // Create a function that takes the variable and returns the expression result
      const fn = new Function(variable, `return (${expression});`);
      result.push(fn(value));
    } catch (e) {
      throw new Error('Invalid expression');
    }
  }
  
  return result;
}