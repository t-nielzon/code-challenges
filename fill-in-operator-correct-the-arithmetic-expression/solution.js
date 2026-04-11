function correct(exp) {
  const [left, right] = exp.split('=').map(s => s.trim());
  const target = parseFloat(right);
  const placeholders = (left.match(/\(\)/g) || []).length;
  const ops = ['+', '-', '*', '/'];
  const total = 4 ** placeholders;

  for (let i = 0; i < total; i++) {
    const combo = [];
    let num = i;
    for (let j = 0; j < placeholders; j++) {
      combo.push(ops[num % 4]);
      num = num >> 2;
    }

    let idx = 0;
    const filled = left.replace(/\(\)/g, () => `(${combo[idx++]})`);
    const evalExpr = filled.replace(/\(([+\-*/])\)/g, '$1');

    try {
      const result = Function('"use strict"; return (' + evalExpr + ')')();
      if (Math.abs(result - target) < 1e-9) {
        return filled + ' = ' + right;
      }
    } catch (e) {
      continue;
    }
  }
}