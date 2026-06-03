function diag_1_sym(s) {
  const rows = s.split("\n");
  const n = rows.length;
  const result = [];
  for (let col = 0; col < n; col++) {
    let line = "";
    for (let row = 0; row < n; row++) {
      line += rows[row][col];
    }
    result.push(line);
  }
  return result.join("\n");
}

function rot_90_clock(s) {
  const rows = s.split("\n");
  const n = rows.length;
  const result = [];
  for (let col = 0; col < n; col++) {
    let line = "";
    for (let row = n - 1; row >= 0; row--) {
      line += rows[row][col];
    }
    result.push(line);
  }
  return result.join("\n");
}

function selfie_and_diag1(s) {
  const rows = s.split("\n");
  const diag = diag_1_sym(s).split("\n");
  return rows.map((row, i) => row + "|" + diag[i]).join("\n");
}

function oper(fct, s) {
  return fct(s);
}