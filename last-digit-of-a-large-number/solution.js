function lastDigit(a, b) {
  const cycles = [
    [0],
    [1],
    [2, 4, 8, 6],
    [3, 9, 7, 1],
    [4, 6],
    [5],
    [6],
    [7, 9, 3, 1],
    [8, 4, 2, 6],
    [9, 1]
  ];

  if (b === 0n) {
    return 1n;
  }

  if (a === 0n) {
    return 0n;
  }

  const lastDigitA = Number(a % 10n);
  const cycle = cycles[lastDigitA];
  const index = Number((b - 1n) % BigInt(cycle.length));
  
  return BigInt(cycle[index]);
}