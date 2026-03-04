function listComprehension(str) {
  str = str.trim();
  if (str[0] !== '[' || str[str.length - 1] !== ']') throw new Error('Invalid');

  var inner = str.slice(1, -1).trim();
  var match = inner.match(/^(.*?)\s+for\s+(\w+)\s+in\s+([\s\S]+)$/);
  if (!match) throw new Error('Invalid');

  var expr = match[1];
  var varName = match[2];
  var source = match[3];

  var arr;
  try {
    arr = new Function('return ' + source)();
  } catch (e) {
    throw new Error('Invalid');
  }

  if (!Array.isArray(arr)) throw new Error('Invalid');

  var fn;
  try {
    fn = new Function(varName, 'return ' + expr);
  } catch (e) {
    throw new Error('Invalid');
  }

  return arr.map(function(val) { return fn(val); });
}