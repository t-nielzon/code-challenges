function diag2Sym(s) {
  const lines = s.split('\n');
  const n = lines.length;
  const result = [];
  for (let i = 0; i < n; i++) {
    let row = '';
    for (let j = 0; j < n; j++) {
      row += lines[n - 1 - j][n - 1 - i];
    }
    result.push(row);
  }
  return result.join('\n');
}

function rot90Counter(s) {
  const lines = s.split('\n');
  const n = lines.length;
  const result = [];
  for (let i = 0; i < n; i++) {
    let row = '';
    for (let j = 0; j < n; j++) {
      row += lines[j][n - 1 - i];
    }
    result.push(row);
  }
  return result.join('\n');
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