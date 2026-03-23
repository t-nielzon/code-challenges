function code(s) {
  const n = Math.ceil(Math.sqrt(s.length));
  const padded = s.padEnd(n * n, String.fromCharCode(11));
  const rows = [];
  for (let i = 0; i < n; i++) {
    rows.push(padded.slice(i * n, (i + 1) * n));
  }
  // clockwise 90°: new[col][n-1-row] = old[row][col]
  const result = [];
  for (let col = 0; col < n; col++) {
    let line = '';
    for (let row = n - 1; row >= 0; row--) {
      line += rows[row][col];
    }
    result.push(line);
  }
  return result.join('\n');
}

function decode(s) {
  const rows = s.split('\n');
  const n = rows.length;
  // counter-clockwise 90°: new[n-1-col][row] = old[row][col]
  const result = [];
  for (let col = n - 1; col >= 0; col--) {
    let line = '';
    for (let row = 0; row < n; row++) {
      line += rows[row][col];
    }
    result.push(line);
  }
  return result.join('').replace(/\x0B+$/, '');
}