function solve(str) {
  function isPalindrome(s) {
    return s === s.split('').reverse().join('');
  }
  
  if (isPalindrome(str)) {
    return "OK";
  }
  
  for (let i = 0; i < str.length; i++) {
    const newStr = str.slice(0, i) + str.slice(i + 1);
    if (isPalindrome(newStr)) {
      return "remove one";
    }
  }
  
  return "not possible";
}