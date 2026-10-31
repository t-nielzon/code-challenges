function numberOfCoins(tips) {
  // Calculate M = product of all moduli using Chinese Remainder Theorem
  let M = 1n;
  for (const [_, m] of tips) {
    M *= BigInt(m);
  }
  
  let result = 0n;
  
  // For each tip, calculate contribution using CRT
  for (const [a, m] of tips) {
    const a_big = BigInt(a);
    const m_big = BigInt(m);
    const M_i = M / m_big;
    
    // Find modular multiplicative inverse of M_i mod m_i
    const y_i = modInverse(M_i, m_big);
    
    result += a_big * M_i * y_i;
  }
  
  return result % M;
}

function modInverse(a, m) {
  // Extended Euclidean Algorithm to find modular multiplicative inverse
  let [old_r, r] = [a, m];
  let [old_s, s] = [1n, 0n];
  
  while (r !== 0n) {
    const quotient = old_r / r;
    [old_r, r] = [r, old_r - quotient * r];
    [old_s, s] = [s, old_s - quotient * s];
  }
  
  if (old_s < 0n) old_s += m;
  
  return old_s;
}