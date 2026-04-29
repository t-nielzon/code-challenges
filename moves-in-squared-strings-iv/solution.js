function diag2Sym(s) {
  const a = s.split('\n');
  const n = a.length;
  const r = [];
  for (let i = 0; i < n; i++) {
    let row = '';
    for (let j = 0; j < n; j++) {
      row += a[n - 1 - j][n - 1 - i];
    }
    r.push(row);
  }
  return r.join('\n');
}

function rot90Counter(s) {
  const a = s.split('\n');
  const n = a.length;
  const r = [];
  for (let i = 0; i < n; i++) {
    let row = '';
    for (let j = 0; j < n; j++) {
      row += a[j][n - 1 - i];
    }
    r.push(row);
  }
  return r.join('\n');
}

function selfieDiag2Counterclock(s) {
  const a = s.split('\n');
  const b = diag2Sym(s).split('\n');
  const c = rot90Counter(s).split('\n');
  return a.map((row, i) => `${row}|${b[i]}|${c[i]}`).join('\n');
}

function oper(fct, s) {
  return fct(s);
}