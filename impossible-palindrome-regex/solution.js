const palindromeMatcher = (() => {
  const isPalindrome = (word) => {
    const s = word.toLowerCase();
    for (let i = 0, j = s.length - 1; i < j; i++, j--) {
      if (s[i] !== s[j]) return false;
    }
    return true;
  };

  // the well-known string methods delegate to the matcher's symbol methods,
  // so for each input string we can build a plain regex that matches exactly
  // the palindromic words found in that specific string, then delegate to it
  const buildRegex = (input) => {
    const str = String(input);
    const words = str.match(/[a-zA-Z]+/g) || [];
    const pals = [...new Set(words.filter(isPalindrome))];
    if (pals.length === 0) return /(?!)/g; // never matches
    return new RegExp(`\\b(?:${pals.join('|')})\\b`, 'g');
  };

  return {
    [Symbol.match](str) {
      return buildRegex(str)[Symbol.match](str);
    },
    [Symbol.matchAll](str) {
      return buildRegex(str)[Symbol.matchAll](str);
    },
    [Symbol.replace](str, replacement) {
      return buildRegex(str)[Symbol.replace](str, replacement);
    },
    [Symbol.search](str) {
      return buildRegex(str)[Symbol.search](str);
    },
    [Symbol.split](str, limit) {
      return buildRegex(str)[Symbol.split](str, limit);
    },
  };
})();