function diag_1_sym(s) {
  const lines = s.split('\n');
  const n = lines.length;
  const result = [];
  
  for (let j = 0; j < n; j++) {
    let row = '';
    for (let i = 0; i < n; i++) {
      row += lines[i][j];
    }
    result.push(row);
  }
  
  return result.join('\n');
}

function rot_90_clock(s) {
  const lines = s.split('\n');
  const n = lines.length;
  const result = [];
  
  for (let i = 0; i < n; i++) {
    let row = '';
    for (let j = n - 1; j >= 0; j--) {
      row += lines[j][i];
    }
    result.push(row);
  }
  
  return result.join('\n');
}

function selfie_and_diag1(s) {
  const diag = diag_1_sym(s);
  const lines = s.split('\n');
  const diagLines = diag.split('\n');
  const result = [];
  
  for (let i = 0; i < lines.length; i++) {
    result.push(lines[i] + '|' + diagLines[i]);
  }
  
  return result.join('\n');
}

function oper(fct, s) {
  return fct(s);
}