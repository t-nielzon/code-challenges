function code(t) {
  const l = t.length;
  const n = Math.ceil(Math.sqrt(l));
  
  const padChar = String.fromCharCode(11);
  const padded = t + padChar.repeat(n * n - l);
  
  const lines = [];
  for (let i = 0; i < n; i++) {
    let row = '';
    for (let j = 0; j < n; j++) {
      row += padded[(n - 1 - j) * n + i];
    }
    lines.push(row);
  }
  
  return lines.join('\n');
}

function decode(s) {
  const lines = s.split('\n');
  const n = lines.length;
  
  let text = '';
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
      text += lines[j][n - 1 - i];
    }
  }
  
  const padChar = String.fromCharCode(11);
  return text.replace(new RegExp(padChar + '+$'), '');
}