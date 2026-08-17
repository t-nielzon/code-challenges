function computeSequences(n) {
  const a = [1];
  const j = [0];
  
  for (let day = 1; day < n; day++) {
    const prev_a = a[day - 1];
    const prev_j = j[day - 1];
    
    a.push(day - (j[prev_a] || 0));
    j.push(day - (a[prev_j] || 0));
  }
  
  return { a, j };
}

function ann(n) {
  return computeSequences(n).a;
}

function john(n) {
  return computeSequences(n).j;
}

function sum_ann(n) {
  const { a } = computeSequences(n);
  return a.reduce((sum, val) => sum + val, 0);
}

function sum_john(n) {
  const { j } = computeSequences(n);
  return j.reduce((sum, val) => sum + val, 0);
}