function isPalindrome(n) {
  const s = '' + n;
  for (let a = 0, b = s.length - 1; a < b; a++, b--) {
    if (s[a] !== s[b]) return false;
  }
  return true;
}

function largestPalindrome(lower, upper) {
  let best = NaN;
  for (let i = upper; i >= lower; i--) {
    // largest possible product using i is i * upper; if that can't beat best,
    // no smaller i can either, so we're done
    if (i * upper <= best) break;
    for (let j = upper; j >= i; j--) {
      const product = i * j;
      // products shrink as j shrinks, so once we drop to best, stop this row
      if (product <= best) break;
      if (isPalindrome(product)) {
        best = product;
        break;
      }
    }
  }
  return best;
}