function correct(exp) {
  const [left, right] = exp.split(' = ');
  const target = Number(right);
  const ops = ['+', '-', '*', '/'];
  const slots = (left.match(/\(\)/g) || []).length;

  function tryOps(idx, current) {
    if (idx === slots) {
      let evalStr = current.replace(/\(([+\-*/])\)/g, ' $1 ');
      try {
        if (Math.abs(Function('"use strict"; return (' + evalStr + ')')() - target) < 1e-9) {
          return current;
        }
      } catch (e) {}
      return null;
    }
    for (const op of ops) {
      const next = current.replace('()', '(' + op + ')');
      const result = tryOps(idx + 1, next);
      if (result) return result;
    }
    return null;
  }

  return tryOps(0, left) + ' = ' + right;
}