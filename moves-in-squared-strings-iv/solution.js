function diag_2_sym(s) {
  const rows = s.split("\n");
  const n = rows.length;
  const out = [];
  for (let i = n - 1; i >= 0; i--) {
    let line = "";
    for (let j = n - 1; j >= 0; j--) {
      line += rows[j][i];
    }
    out.push(line);
  }
  return out.join("\n");
}

function rot_90_counter(s) {
  const rows = s.split("\n");
  const n = rows.length;
  const out = [];
  for (let i = n - 1; i >= 0; i--) {
    let line = "";
    for (let j = 0; j < n; j++) {
      line += rows[j][i];
    }
    out.push(line);
  }
  return out.join("\n");
}

function selfie_diag2_counterclock(s) {
  const a = s.split("\n");
  const b = diag_2_sym(s).split("\n");
  const c = rot_90_counter(s).split("\n");
  return a.map((row, i) => row + "|" + b[i] + "|" + c[i]).join("\n");
}

function oper(fct, s) {
  return fct(s);
}