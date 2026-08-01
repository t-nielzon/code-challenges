const cache = {};

function getSequences(n) {
  if (cache[n]) return cache[n];
  
  let a = [1], j = [0];
  for (let day = 1; day < n; day++) {
    let t_j = j[day - 1];
    j.push(day - a[t_j]);
    
    let t_a = a[day - 1];
    a.push(day - j[t_a]);
  }
  cache[n] = { a, j };
  return { a, j };
}

function ann(n) {
  if (n === 0) return [];
  return getSequences(n).a;
}

function john(n) {
  if (n === 0) return [];
  return getSequences(n).j;
}

function sum_ann(n) {
  if (n === 0) return 0;
  let { a } = getSequences(n);
  return a.reduce((sum, x) => sum + x, 0);
}

function sum_john(n) {
  if (n === 0) return 0;
  let { j } = getSequences(n);
  return j.reduce((sum, x) => sum + x, 0);
}