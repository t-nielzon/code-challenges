function cubicNumbers(s) {
  const runs = s.match(/\d+/g) || [];
  const groups = [];
  for (const run of runs) {
    for (let i = 0; i < run.length; i += 3) {
      groups.push(run.slice(i, i + 3));
    }
  }
  const cubic = [];
  let sum = 0;
  for (const g of groups) {
    const n = parseInt(g, 10);
    const cubeSum = [...g].reduce((acc, d) => acc + Math.pow(parseInt(d, 10), 3), 0);
    if (cubeSum === n) {
      cubic.push(n);
      sum += n;
    }
  }
  if (cubic.length === 0) return "Unlucky";
  return `${cubic.join(' ')} ${sum} Lucky`;
}