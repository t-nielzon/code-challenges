const palindromeMatcher = {
  [Symbol.match](str) {
    const regex = /\b[a-z]+\b/gi;
    const matches = [];
    let match;
    
    while ((match = regex.exec(str)) !== null) {
      const word = match[0];
      if (word.toLowerCase() === word.toLowerCase().split('').reverse().join('')) {
        matches.push(word);
      }
    }
    
    return matches.length > 0 ? matches : null;
  },
  
  [Symbol.matchAll](str) {
    const regex = /\b[a-z]+\b/gi;
    const matches = [];
    let match;
    
    while ((match = regex.exec(str)) !== null) {
      const word = match[0];
      if (word.toLowerCase() === word.toLowerCase().split('').reverse().join('')) {
        matches.push(match);
      }
    }
    
    return matches[Symbol.iterator]();
  },
  
  [Symbol.replace](str, replacement) {
    const regex = /\b[a-z]+\b/gi;
    const matches = [];
    let match;
    
    while ((match = regex.exec(str)) !== null) {
      const word = match[0];
      if (word.toLowerCase() === word.toLowerCase().split('').reverse().join('')) {
        matches.push(match);
      }
    }
    
    if (matches.length === 0) return str;
    
    let result = '';
    let lastIndex = 0;
    
    for (const m of matches) {
      result += str.substring(lastIndex, m.index);
      const repl = replacement.replace(/\$&/g, m[0]);
      result += repl;
      lastIndex = m.index + m[0].length;
    }
    result += str.substring(lastIndex);
    
    return result;
  },
  
  [Symbol.search](str) {
    const regex = /\b[a-z]+\b/gi;
    let match;
    
    while ((match = regex.exec(str)) !== null) {
      const word = match[0];
      if (word.toLowerCase() === word.toLowerCase().split('').reverse().join('')) {
        return match.index;
      }
    }
    return -1;
  },
  
  [Symbol.split](str, limit) {
    const regex = /\b[a-z]+\b/gi;
    const matches = [];
    let match;
    
    while ((match = regex.exec(str)) !== null) {
      const word = match[0];
      if (word.toLowerCase() === word.toLowerCase().split('').reverse().join('')) {
        matches.push(match);
      }
    }
    
    if (matches.length === 0) return [str];
    
    const result = [];
    let lastIndex = 0;
    for (const m of matches) {
      result.push(str.substring(lastIndex, m.index));
      lastIndex = m.index + m[0].length;
    }
    result.push(str.substring(lastIndex));
    
    return limit ? result.slice(0, limit) : result;
  }
};