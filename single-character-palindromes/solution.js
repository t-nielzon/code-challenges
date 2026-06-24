function isPalindrome(s, i, j) {
  while (i < j) {
    if (s[i] !== s[j]) return false;
    i++;
    j--;
  }
  return true;
}

function solve(s) {
  let i = 0;
  let j = s.length - 1;

  while (i < j) {
    if (s[i] !== s[j]) {
      // mismatch found: try removing one of the two mismatched characters
      if (isPalindrome(s, i + 1, j) || isPalindrome(s, i, j - 1)) {
        return "remove one";
      }
      return "not possible";
    }
    i++;
    j--;
  }

  return "OK";
}