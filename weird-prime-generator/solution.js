function gcd(a, b) {
  while (b) {
    [a, b] = [b, a % b];
  }
  return a;
}

function an(n) {
  if (n === 0) return [];
  const result = [7];
  for (let i = 2; i <= n; i++) {
    const prev = result[result.length - 1];
    result.push(prev + gcd(i, prev));
  }
  return result;
}

function gn(n) {
  if (n === 0) return [];
  const aSeq = an(n);
  const result = [1];
  for (let i = 1; i < aSeq.length; i++) {
    result.push(aSeq[i] - aSeq[i - 1]);
  }
  return result;
}

function countOnes(n) {
  const gSeq = gn(n);
  return gSeq.filter(x => x === 1).length;
}

function pn(n) {
  const result = [];
  const seen = new Set();
  let i = 1;
  
  while (result.length < n) {
    const gSeq = gn(i);
    if (gSeq[i - 1] !== 1 && !seen.has(gSeq[i - 1])) {
      result.push(gSeq[i - 1]);
      seen.add(gSeq[i - 1]);
    }
    i++;
  }
  
  return result;
}

function maxPn(n) {
  const primes = pn(n);
  return Math.max(...primes);
}

function anOver(n) {
  const result = [];
  let i = 1;
  
  while (result.length < n) {
    const gSeq = gn(i);
    if (gSeq[i - 1] !== 1) {
      const aSeq = an(i);
      result.push(aSeq[i - 1] / i);
    }
    i++;
  }
  
  return result;
}

function anOverAverage(n) {
  const vals = anOver(n);
  if (vals.length === 0) return 0;
  const sum = vals.reduce((a, b) => a + b, 0);
  return Math.floor(sum / vals.length);
}