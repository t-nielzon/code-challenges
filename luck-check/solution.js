function luckCheck(str) {
  if (typeof str !== 'string' || str.length === 0 || !/^\d+$/.test(str)) {
    throw new Error('Invalid input');
  }
  const half = Math.floor(str.length / 2);
  const left = str.slice(0, half);
  const right = str.slice(str.length - half);
  const sum = s => [...s].reduce((a, c) => a + Number(c), 0);
  return sum(left) === sum(right);
}