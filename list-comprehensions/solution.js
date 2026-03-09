function listComprehension(str) {
  str = str.trim();

  if (str[0] !== '[' || str[str.length - 1] !== ']') {
    throw new Error('Invalid list comprehension');
  }

  var inner = str.slice(1, -1).trim();

  // split on " for " to separate expression from iteration
  var forIndex = inner.indexOf(' for ');
  if (forIndex === -1) throw new Error('Invalid list comprehension');

  var expr = inner.slice(0, forIndex).trim();
  var rest = inner.slice(forIndex + 5).trim();

  // split on " in " to separate variable from iterable
  var inIndex = rest.indexOf(' in ');
  if (inIndex === -1) throw new Error('Invalid list comprehension');

  var varName = rest.slice(0, inIndex).trim();
  var iterableExpr = rest.slice(inIndex + 4).trim();

  if (!varName || !expr || !iterableExpr) {
    throw new Error('Invalid list comprehension');
  }

  var arr;
  try {
    arr = Function('return (' + iterableExpr + ')')();
  } catch (e) {
    throw new Error('Invalid list comprehension');
  }

  if (!Array.isArray(arr)) throw new Error('Invalid list comprehension');

  var mapFn;
  try {
    mapFn = Function(varName, 'return (' + expr + ')');
  } catch (e) {
    throw new Error('Invalid list comprehension');
  }

  return arr.map(mapFn);
}

module.exports = listComprehension;