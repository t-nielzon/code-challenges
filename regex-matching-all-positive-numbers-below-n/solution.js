function regexBelowN(n) {
  const s = String(n);
  const L = s.length;
  const parts = [];

  for (let d = 1; d < L; d++) {
    if (d === 1) parts.push('[1-9]');
    else parts.push('[1-9]\\d{' + (d - 1) + '}');
  }

  for (let i = 0; i < L; i++) {
    const prefix = s.slice(0, i);
    const di = parseInt(s[i], 10);
    const lo = i === 0 ? 1 : 0;
    const hi = di - 1;
    if (lo > hi) continue;

    const charClass = lo === hi ? String(lo) : '[' + lo + '-' + hi + ']';
    const suffixLen = L - 1 - i;
    let suffix = '';
    if (suffixLen === 1) suffix = '\\d';
    else if (suffixLen > 1) suffix = '\\d{' + suffixLen + '}';

    parts.push(prefix + charClass + suffix);
  }

  if (parts.length === 0) return '^(?!)$';
  return '^(?:' + parts.join('|') + ')$';
}