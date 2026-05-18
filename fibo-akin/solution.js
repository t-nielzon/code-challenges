function buildU(n) {
  const u = new Array(n + 1);
  u[1] = 1;
  if (n >= 2) u[2] = 1;
  for (let i = 3; i <= n; i++) {
    u[i] = u[i - u[i - 1]] + u[i - u[i - 2]];
  }
  return u;
}

function lengthSupUK(n, k) {
  const u = buildU(n);
  let count = 0;
  for (let i = 1; i <= n; i++) {
    if (u[i] >= k) count++;
  }
  return count;
}

function comp(n) {
  const u = buildU(n);
  let count = 0;
  for (let i = 2; i <= n; i++) {
    if (u[i] < u[i - 1]) count++;
  }
  return count;
}