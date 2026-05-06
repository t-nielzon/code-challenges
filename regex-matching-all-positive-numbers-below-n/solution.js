function regexAllNumbersLessThan(n) {
  const m = BigInt(n) - 1n;
  if (m < 1n) return '^(?!)$';

  const s = m.toString();
  const L = s.length;
  const parts = [];

  for (let i = 1; i < L; i++) {
    parts.push(i === 1 ? '[1-9]' : `[1-9]\\d{${i - 1}}`);
  }

  const d1 = parseInt(s[0], 10);
  if (L === 1) {
    if (d1 > 1) parts.push(d1 === 2 ? '1' : `[1-${d1 - 1}]`);
    parts.push(s);
  } else {
    if (d1 > 1) {
      const range = d1 === 2 ? '1' : `[1-${d1 - 1}]`;
      parts.push(`${range}\\d{${L - 1}}`);
    }
    parts.push(s[0] + leqRegex(s.slice(1)));
  }

  return '^(?:' + parts.join('|') + ')$';
}

function leqRegex(s) {
  if (s.length === 0) return '';
  const d = parseInt(s[0], 10);
  const rest = s.slice(1);
  const alts = [];

  if (d > 0) {
    const range = d === 1 ? '0' : `[0-${d - 1}]`;
    alts.push(rest.length === 0 ? range : `${range}\\d{${rest.length}}`);
  }
  alts.push(rest.length === 0 ? s[0] : s[0] + leqRegex(rest));

  return alts.length === 1 ? alts[0] : '(?:' + alts.join('|') + ')';
}