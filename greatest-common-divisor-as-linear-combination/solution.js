// solution.js
function gcd_coeff(a, b) {
  let [old_r, r] = [a, b];
  let [old_s, s] = [1, 0];
  let [old_t, t] = [0, 1];
  
  while (r !== 0) {
    const q = Math.floor(old_r / r);
    [old_r, r] = [r, old_r - q * r];
    [old_s, s] = [s, old_s - q * s];
    [old_t, t] = [t, old_t - q * t];
  }
  
  return [old_s, old_t];
}