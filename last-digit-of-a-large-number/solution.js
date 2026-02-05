function lastDigit(a, b) {
  // Handle 0^0 = 1
  if (b === 0n) return 1n;
  
  // Get the last digit of a
  const lastDigitA = a % 10n;
  
  // If last digit is 0, result is always 0
  if (lastDigitA === 0n) return 0n;
  
  // The cycle lengths for each digit 0-9:
  // 0: [0] - cycle length 1
  // 1: [1] - cycle length 1
  // 2: [2,4,8,6] - cycle length 4
  // 3: [3,9,7,1] - cycle length 4
  // 4: [4,6] - cycle length 2
  // 5: [5] - cycle length 1
  // 6: [6] - cycle length 1
  // 7: [7,9,3,1] - cycle length 4
  // 8: [8,4,2,6] - cycle length 4
  // 9: [9,1] - cycle length 2
  
  // All cycles divide evenly into 4, so we can use b % 4
  // But we need to handle when b % 4 === 0 (meaning position 4 in the cycle)
  
  const cycles = {
    0n: [0n],
    1n: [1n],
    2n: [2n, 4n, 8n, 6n],
    3n: [3n, 9n, 7n, 1n],
    4n: [4n, 6n],
    5n: [5n],
    6n: [6n],
    7n: [7n, 9n, 3n, 1n],
    8n: [8n, 4n, 2n, 6n],
    9n: [9n, 1n]
  };
  
  const cycle = cycles[lastDigitA];
  const cycleLen = BigInt(cycle.length);
  
  // Get the position in cycle (1-indexed becomes 0-indexed)
  // b=1 should give index 0, b=2 index 1, etc.
  // b % cycleLen gives 0 when b is multiple of cycleLen, which should be last position
  let pos = b % cycleLen;
  if (pos === 0n) {
    pos = cycleLen;
  }
  
  return cycle[Number(pos - 1n)];
}