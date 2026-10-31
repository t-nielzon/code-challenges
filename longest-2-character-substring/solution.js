function longest2CharSubstring(str) {
  if (str.length === 0) return '';
  
  let maxSubstring = '';
  let left = 0;
  let charCount = {};
  
  for (let right = 0; right < str.length; right++) {
    charCount[str[right]] = (charCount[str[right]] || 0) + 1;
    
    while (Object.keys(charCount).length > 2) {
      charCount[str[left]]--;
      if (charCount[str[left]] === 0) {
        delete charCount[str[left]];
      }
      left++;
    }
    
    const currentSubstring = str.substring(left, right + 1);
    if (currentSubstring.length > maxSubstring.length) {
      maxSubstring = currentSubstring;
    }
  }
  
  return maxSubstring;
}