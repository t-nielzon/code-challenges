function length_sup_u_k(n, k) {
  const u = new Array(n + 1);
  u[1] = 1;
  u[2] = 1;
  
  for (let i = 3; i <= n; i++) {
    u[i] = u[i - u[i - 1]] + u[i - u[i - 2]];
  }
  
  let count = 0;
  for (let i = 1; i <= n; i++) {
    if (u[i] >= k) {
      count++;
    }
  }
  return count;
}

function comp(n) {
  const u = new Array(n + 1);
  u[1] = 1;
  u[2] = 1;
  
  for (let i = 3; i <= n; i++) {
    u[i] = u[i - u[i - 1]] + u[i - u[i - 2]];
  }
  
  let count = 0;
  for (let i = 2; i <= n; i++) {
    if (u[i] < u[i - 1]) {
      count++;
    }
  }
  return count;
}