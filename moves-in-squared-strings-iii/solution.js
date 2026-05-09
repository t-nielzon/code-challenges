function diag1Sym(s) {
  const rows = s.split("\n");
  const n = rows.length;
  const result = [];
  for (let i = 0; i < n; i++) {
    let line = "";
    for (let j = 0; j < n; j++) {
      line += rows[j][i];
    }
    result.push(line);
  }
  return result.join("\n");
}

function rot90Clock(s) {
  const rows = s.split("\n");
  const n = rows.length;
  const result = [];
  for (let i = 0; i < n; i++) {
    let line = "";
    for (let j = n - 1; j >= 0; j--) {
      line += rows[j][i];
    }
    result.push(line);
  }
  return result.join("\n");
}

function selfieAndDiag1(s) {
  const rows = s.split("\n");
  const sym = diag1Sym(s).split("\n");
  return rows.map((row, i) => row + "|" + sym[i]).join("\n");
}

function oper(fct, s) {
  return fct(s);
}