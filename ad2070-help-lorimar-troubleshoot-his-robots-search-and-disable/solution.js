function helpLorimar(log) {
  function isPrime(num) {
    if (num < 2) return false;
    if (num === 2) return true;
    if (num % 2 === 0) return false;
    for (let i = 3; i * i <= num; i += 2) {
      if (num % i === 0) return false;
    }
    return true;
  }

  const counts = {};
  for (let num of log) {
    const str = String(num);
    counts[str] = (counts[str] || 0) + 1;
  }

  let totalCount = 0;
  for (let [numStr, count] of Object.entries(counts)) {
    if (numStr.length !== 4) continue;
    if (numStr[2] !== '2' && numStr[2] !== '3') continue;
    if (count <= 3) continue;
    if (!isPrime(parseInt(numStr))) continue;
    totalCount += count;
  }

  return totalCount > 50 ? "match disable bot" : "no match continue";
}