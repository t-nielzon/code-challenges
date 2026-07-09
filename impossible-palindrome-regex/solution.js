const palindromeMatcher = {
  [Symbol.match](str) {
    const regex = /\b[a-z]+\b/gi;
    let match;
    const palindromes = [];
    while ((match = regex.exec(str)) !== null) {
      const word = match[0];
      const lower = word.toLowerCase();
      if (lower === lower.split('').reverse().join('')) {
        palindromes.push(word);
      }
    }
    return palindromes.length > 0 ? palindromes : null;
  },

  [Symbol.matchAll](str) {
    const regex = /\b[a-z]+\b/gi;
    let match;
    const results = [];
    while ((match = regex.exec(str)) !== null) {
      const word = match[0];
      const lower = word.toLowerCase();
      if (lower === lower.split('').reverse().join('')) {
        const matchObject = [word];
        matchObject.index = match.index;
        matchObject.input = str;
        matchObject.groups = undefined;
        results.push(matchObject);
      }
    }
    return results[Symbol.iterator]();
  },

  [Symbol.replace](str, replaceValue) {
    const regex = /\b[a-z]+\b/gi;
    let match;
    const palindromes = [];
    while ((match = regex.exec(str)) !== null) {
      const word = match[0];
      const lower = word.toLowerCase();
      if (lower === lower.split('').reverse().join('')) {
        palindromes.push({ word, index: match.index });
      }
    }
    
    for (let i = palindromes.length - 1; i >= 0; i--) {
      const { word, index } = palindromes[i];
      let replacement;
      if (typeof replaceValue === 'function') {
        replacement = replaceValue(word, index, str);
      } else {
        replacement = String(replaceValue).replace(/\$&/g, word);
      }
      str = str.slice(0, index) + replacement + str.slice(index + word.length);
    }
    return str;
  },

  [Symbol.search](str) {
    const regex = /\b[a-z]+\b/gi;
    let match;
    while ((match = regex.exec(str)) !== null) {
      const word = match[0];
      const lower = word.toLowerCase();
      if (lower === lower.split('').reverse().join('')) {
        return match.index;
      }
    }
    return -1;
  },

  [Symbol.split](str) {
    const regex = /\b[a-z]+\b/gi;
    let match;
    const parts = [];
    let lastIndex = 0;
    while ((match = regex.exec(str)) !== null) {
      const word = match[0];
      const lower = word.toLowerCase();
      if (lower === lower.split('').reverse().join('')) {
        parts.push(str.slice(lastIndex, match.index));
        lastIndex = match.index + word.length;
      }
    }
    parts.push(str.slice(lastIndex));
    return parts;
  }
};