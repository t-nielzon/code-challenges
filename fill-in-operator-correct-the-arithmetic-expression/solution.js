function correct(exp) {
  const [left, right] = exp.split('=').map(s => s.trim());
  const target = parseFloat(right);
  const slots = (left.match(/\(\)/g) || []).length;
  const ops = ['+', '-', '*', '/'];
  const total = Math.pow(4, slots);

  for (let i = 0; i < total; i++) {
    const combo = [];
    let num = i;
    for (let j = 0; j < slots; j++) {
      combo.push(ops[num % 4]);
      num = Math.floor(num / 4);
    }

    let idx = 0;
    const filled = left.replace(/\(\)/g, () => `(${combo[idx++]})`);
    const evalStr = filled.replace(/\(([+\-*/])\)/g, '$1');

    try {
      const result = Function('"use strict"; return (' + evalStr + ')')();
      if (Math.abs(result - target) < 1e-9) {
        return filled + ' = ' + right;
      }
    } catch (e) {}
  }
}