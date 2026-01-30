function stringTransformation(s, pattern) {
  let result = s;
  
  for (const op of pattern) {
    switch (op) {
      case 'U':
        result = result.toUpperCase();
        break;
      case 'L':
        result = result.toLowerCase();
        break;
      case 'R':
        result = result.split('').reverse().join('');
        break;
      case 'D':
        result = result.split('').map(c => c + c).join('');
        break;
    }
  }
  
  return result;
}