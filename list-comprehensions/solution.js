function listComprehension(str) {
  // Trim whitespace
  str = str.trim();
  
  // Check if it starts with [ and ends with ]
  if (!str.startsWith('[') || !str.endsWith(']')) {
    throw new Error('Invalid list comprehension');
  }
  
  // Remove outer brackets
  const inner = str.slice(1, -1).trim();
  
  // Parse the list comprehension pattern: expression for variable in iterable
  const forMatch = inner.match(/^(.+?)\s+for\s+(\w+)\s+in\s+(.+)$/);
  
  if (!forMatch) {
    throw new Error('Invalid list comprehension');
  }
  
  const expression = forMatch[1].trim();
  const variable = forMatch[2].trim();
  const iterableExpr = forMatch[3].trim();
  
  // Validate variable name
  if (!/^[a-zA-Z_$][a-zA-Z0-9_$]*$/.test(variable)) {
    throw new Error('Invalid variable name');
  }
  
  // Evaluate the iterable expression
  let iterable;
  try {
    iterable = eval(iterableExpr);
  } catch (e) {
    throw new Error('Invalid iterable expression');
  }
  
  // Check if iterable is actually iterable (array-like)
  if (!Array.isArray(iterable)) {
    throw new Error('Iterable must be an array');
  }
  
  // Build result by evaluating expression for each element
  const result = [];
  for (let i = 0; i < iterable.length; i++) {
    const value = iterable[i];
    try {
      // Create a function that evaluates the expression with the variable bound
      const fn = new Function(variable, `return (${expression});`);
      result.push(fn(value));
    } catch (e) {
      throw new Error('Invalid expression');
    }
  }
  
  return result;
}