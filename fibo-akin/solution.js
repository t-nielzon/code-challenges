function length_sup_u_k(n, k) {
  const seq = [0, 1, 1];
  let count = 0;
  
  if (seq[1] >= k) count++;
  if (seq[2] >= k) count++;
  
  for (let i = 3; i <= n; i++) {
    seq[i] = seq[i - seq[i - 1]] + seq[i - seq[i - 2]];
    if (seq[i] >= k) count++;
  }
  
  return count;
}

function comp(n) {
  const seq = [0, 1, 1];
  let count = 0;
  
  for (let i = 3; i <= n; i++) {
    seq[i] = seq[i - seq[i - 1]] + seq[i - seq[i - 2]];
    if (seq[i] < seq[i - 1]) count++;
  }
  
  return count;
}