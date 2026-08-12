const palindromeMatcher = {
  [Symbol.match](str) {
    const matches = [];
    const regex = /\b[a-z]+\b/gi;
    let match;
    while ((match = regex.exec(str)) !== null) {
      if (isPalindrome(match[0])) {
        matches.push(match[0]);
      }
    }
    return matches.length > 0 ? matches : null;
  },

  [Symbol.matchAll](str) {
    const matches = [];
    const regex = /\b[a-z]+\b/gi;
    let match;
    while ((match = regex.exec(str)) !== null) {
      if (isPalindrome(match[0])) {
        matches.push(Object.assign([], match, {
          index: match.index,
          input: match.input,
          groups: match.groups
        }));
      }
    }
    return matches[Symbol.iterator]();
  },

  [Symbol.replace](str, replacement) {
    const regex = /\b[a-z]+\b/gi;
    const matches = [];
    let match;
    
    while ((match = regex.exec(str)) !== null) {
      if (isPalindrome(match[0])) {
        matches.push({
          matched: match[0],
          index: match.index
        });
      }
    }
    
    let result = str;
    for (let i = matches.length - 1; i >= 0; i--) {
      const m = matches[i];
      let replaceStr;
      
      if (typeof replacement === 'function') {
        replaceStr = replacement(m.matched, 0, str);
      } else {
        replaceStr = replacement
          .replace(/\$&/g, m.matched)
          .replace(/\$`/g, str.substring(0, m.index))
          .replace(/\$'/g, str.substring(m.index + m.matched.length));
      }
      
      result = result.substring(0, m.index) + replaceStr + result.substring(m.index + m.matched.length);
    }
    
    return result;
  },

  [Symbol.search](str) {
    const regex = /\b[a-z]+\b/gi;
    let match;
    while ((match = regex.exec(str)) !== null) {
      if (isPalindrome(match[0])) {
        return match.index;
      }
    }
    return -1;
  },

  [Symbol.split](str, limit) {
    const regex = /\b[a-z]+\b/gi;
    const result = [];
    let lastIndex = 0;
    let match;
    
    while ((match = regex.exec(str)) !== null) {
      if (isPalindrome(match[0])) {
        result.push(str.substring(lastIndex, match.index));
        lastIndex = match.index + match[0].length;
      }
    }
    
    result.push(str.substring(lastIndex));
    
    if (limit !== undefined) {
      result.length = limit;
    }
    
    return result;
  }
};

function isPalindrome(word) {
  const lower = word.toLowerCase();
  for (let i = 0; i < lower.length / 2; i++) {
    if (lower[i] !== lower[lower.length - 1 - i]) {
      return false;
    }
  }
  return true;
}