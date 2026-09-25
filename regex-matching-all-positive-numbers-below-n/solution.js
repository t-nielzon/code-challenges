function regexRange(n) {
  if (n === 1) {
    return '(?!)';
  }
  
  const max = String(n - 1);
  const parts = [];
  
  for (let len = 1; len < max.length; len++) {
    if (len === 1) {
      parts.push('[1-9]');
    } else {
      parts.push('[1-9][0-9]{' + (len - 1) + '}');
    }
  }
  
  parts.push(rangePattern('', max));
  
  return parts.join('|');
}

function rangePattern(prefix, max) {
  if (prefix.length === max.length) {
    return prefix;
  }
  
  const pos = prefix.length;
  const digit = parseInt(max[pos]);
  const remaining = max.length - pos - 1;
  
  const parts = [];
  const start = pos === 0 ? 1 : 0;
  
  for (let d = start; d < digit; d++) {
    const rest = remaining > 0 ? '[0-9]{' + remaining + '}' : '';
    parts.push(prefix + d + rest);
  }
  
  parts.push(rangePattern(prefix + digit, max));
  
  return parts.length === 1 ? parts[0] : '(' + parts.join('|') + ')';
}