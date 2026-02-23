function lookAndWrite(n) {
  const counts = new Array(10).fill(0);
  counts[1] = 2;
  let len = 2;

  for (let i = 0; i < n; i++) {
    const snap = counts.slice();

    for (let d = 0; d < 10; d++) {
      const c = snap[d];
      if (c === 0) continue;

      // decompose the count into its digits and tally them
      let num = c;
      let numDigits = 0;
      while (num > 0) {
        counts[num % 10]++;
        num = Math.floor(num / 10);
        numDigits++;
      }
      counts[d]++;
      len += numDigits + 1;
    }
  }

  return len;
}