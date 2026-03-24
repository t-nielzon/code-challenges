function belowN(n) {
  if (n <= 1) return '\\b\\B';

  const s = n.toString();
  const len = s.length;
  const digits = s.split('').map(Number);
  const parts = [];

  // match all positive numbers with fewer digits than n
  if (len === 2) {
    parts.push('[1-9]');
  } else if (len > 2) {
    parts.push('[1-9][0-9]{0,' + (len - 2) + '}');
  }

  // match numbers with exactly len digits but numerically < n
  for (let i = 0; i < len; i++) {
    const lo = i === 0 ? 1 : 0;
    const hi = digits[i] - 1;

    if (lo > hi) continue;

    let part = '';
    for (let j = 0; j < i; j++) part += digits[j];

    if (lo === hi) part += lo;
    else part += '[' + lo + '-' + hi + ']';

    const remaining = len - i - 1;
    if (remaining === 1) part += '[0-9]';
    else if (remaining > 1) part += '[0-9]{' + remaining + '}';

    parts.push(part);
  }

  return '^(' + parts.join('|') + ')$';
}