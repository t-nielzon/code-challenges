function solve(a, b) {
  let str = '';
  let num = 2;

  while (str.length < a + b) {
    if (isPrime(num)) str += num;
    num++;
  }

  return str.slice(a, a + b);
}

function isPrime(n) {
  if (n < 2) return false;
  if (n < 4) return true;
  if (n % 2 === 0 || n % 3 === 0) return false;
  for (let i = 5; i * i <= n; i += 6) {
    if (n % i === 0 || n % (i + 2) === 0) return false;
  }
  return true;
}