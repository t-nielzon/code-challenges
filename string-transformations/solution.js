function stringTransformation(s, pattern) {
  // optimize by collapsing consecutive case ops and tracking reverse parity
  // to avoid repeated full-string operations on potentially huge strings
  let reversed = false;
  let caseOp = null; // 'U', 'L', or null
  let dupCount = 0;

  for (const op of pattern) {
    if (op === 'U') {
      caseOp = 'U';
    } else if (op === 'L') {
      caseOp = 'L';
    } else if (op === 'R') {
      reversed = !reversed;
    } else if (op === 'D') {
      dupCount++;
    }
  }

  // apply case first (cheap on original string before duplication)
  if (caseOp === 'U') s = s.toUpperCase();
  else if (caseOp === 'L') s = s.toLowerCase();

  // apply duplications
  if (dupCount > 0) {
    const factor = 2 ** dupCount;
    const parts = [];
    for (let i = 0; i < s.length; i++) {
      parts.push(s[i].repeat(factor));
    }
    s = parts.join('');
  }

  if (reversed) {
    s = s.split('').reverse().join('');
  }

  return s;
}

module.exports = { stringTransformation };