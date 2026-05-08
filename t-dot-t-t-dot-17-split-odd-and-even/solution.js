function splitOddAndEven(n) {
  const digits = String(n).split('');
  const groups = [];
  let current = '';
  let currentParity = null;

  for (const d of digits) {
    const parity = Number(d) % 2;
    if (parity === currentParity) {
      current += d;
    } else {
      if (current !== '') groups.push(current);
      current = d;
      currentParity = parity;
    }
  }
  if (current !== '') groups.push(current);

  return groups.map(Number);
}