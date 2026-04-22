function code(s) {
  const l = s.length;
  const n = Math.ceil(Math.sqrt(l));
  const pad = String.fromCharCode(11);
  let t = s;
  while (t.length < n * n) t += pad;
  const grid = [];
  for (let i = 0; i < n; i++) grid.push(t.slice(i * n, (i + 1) * n));
  const rotated = [];
  for (let i = 0; i < n; i++) {
    let row = '';
    for (let j = n - 1; j >= 0; j--) row += grid[j][i];
    rotated.push(row);
  }
  return rotated.join('\n');
}

function decode(s) {
  const rows = s.split('\n');
  const n = rows.length;
  const rotated = [];
  for (let i = 0; i < n; i++) {
    let row = '';
    for (let j = 0; j < n; j++) row += rows[j][n - 1 - i];
    rotated.push(row);
  }
  const joined = rotated.join('');
  return joined.replace(new RegExp(String.fromCharCode(11) + '+$'), '');
}