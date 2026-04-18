function isPrime(n) {
  if (n < 2) return false;
  if (n < 4) return true;
  if (n % 2 === 0) return false;
  if (n % 3 === 0) return n === 3;
  for (let i = 5; i * i <= n; i += 6) {
    if (n % i === 0 || n % (i + 2) === 0) return false;
  }
  return true;
}

function reverseNumber(n) {
  let r = 0;
  while (n > 0) {
    r = r * 10 + (n % 10);
    n = Math.floor(n / 10);
  }
  return r;
}

function backwardsPrime(start, end) {
  const result = [];
  if (start < 2) start = 2;
  for (let n = start; n <= end; n++) {
    if (n % 10 === 0) continue;
    if (!isPrime(n)) continue;
    const rev = reverseNumber(n);
    if (rev === n) continue;
    if (isPrime(rev)) result.push(n);
  }
  return result;
}