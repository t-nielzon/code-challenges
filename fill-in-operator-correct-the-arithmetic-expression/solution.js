function correct(exp) {
  const [lhs, rhs] = exp.split('=').map(s => s.trim());
  const target = Number(rhs);
  const ops = ['+', '-', '*', '/'];
  const slots = [];
  for (let i = 0; i < lhs.length; i++) {
    if (lhs[i] === '(' && lhs[i + 1] === ')') slots.push(i);
  }
  const n = slots.length;
  const total = Math.pow(4, n);
  for (let m = 0; m < total; m++) {
    const chosen = [];
    let k = m;
    for (let i = 0; i < n; i++) {
      chosen.push(ops[k % 4]);
      k = Math.floor(k / 4);
    }
    let candidate = '';
    let last = 0;
    for (let i = 0; i < n; i++) {
      candidate += lhs.slice(last, slots[i]) + '(' + chosen[i] + ')';
      last = slots[i] + 2;
    }
    candidate += lhs.slice(last);
    let val;
    try {
      val = eval(candidate.replace(/\(([+\-*/])\)/g, '$1'));
    } catch (e) {
      continue;
    }
    if (Math.abs(val - target) < 1e-9) {
      return candidate + ' = ' + rhs;
    }
  }
  return exp;
}