function findX(n) {
  const MOD = 1000000007n;
  if (n === 0) return 0;
  
  const inv2 = 500000004n; // modular inverse of 2
  let sum = 0n;
  let f = 0n;
  
  for (let i = 1; i <= n; i++) {
    const i_n = BigInt(i);
    const term = (3n * i_n % MOD * ((i_n + 1n) % MOD) % MOD * inv2) % MOD;
    f = (sum + term) % MOD;
    sum = (sum + f) % MOD;
  }
  
  return Number(f);
}