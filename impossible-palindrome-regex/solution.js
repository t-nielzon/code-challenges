// the string methods match/matchAll/replace/search/split all delegate to the
// well-known symbols on their argument, so a plain object implementing those
// symbols can "match" palindromes without a real (impossible) regex

const findPalindromes = (str) => {
  const found = [];
  const wordRe = /[a-zA-Z]+/g;
  let m;
  while ((m = wordRe.exec(str)) !== null) {
    const w = m[0].toLowerCase();
    let isPal = true;
    for (let i = 0, j = w.length - 1; i < j; i++, j--) {
      if (w[i] !== w[j]) {
        isPal = false;
        break;
      }
    }
    if (isPal) found.push({ word: m[0], index: m.index });
  }
  return found;
};

// supports the standard $$, $&, $` and $' substitution patterns; there are no
// capture groups, so $n sequences stay literal as the spec dictates
const substitute = (matched, str, position, replacement) =>
  replacement.replace(/\$(\$|&|`|')/g, (_, ch) => {
    if (ch === '$') return '$';
    if (ch === '&') return matched;
    if (ch === '`') return str.slice(0, position);
    return str.slice(position + matched.length);
  });

const palindromeMatcher = {
  // matchAll insists on a global flag once it sees Symbol.match present
  flags: 'g',

  [Symbol.match](str) {
    const s = String(str);
    const matches = findPalindromes(s);
    return matches.length ? matches.map((m) => m.word) : null;
  },

  [Symbol.matchAll](str) {
    const s = String(str);
    const matches = findPalindromes(s);
    return (function* () {
      for (const m of matches) {
        const entry = [m.word];
        entry.index = m.index;
        entry.input = s;
        entry.groups = undefined;
        yield entry;
      }
    })();
  },

  [Symbol.replace](str, replacement) {
    const s = String(str);
    const matches = findPalindromes(s);
    let result = '';
    let last = 0;
    for (const m of matches) {
      result += s.slice(last, m.index);
      result +=
        typeof replacement === 'function'
          ? String(replacement(m.word, m.index, s))
          : substitute(m.word, s, m.index, String(replacement));
      last = m.index + m.word.length;
    }
    return result + s.slice(last);
  },

  [Symbol.search](str) {
    const matches = findPalindromes(String(str));
    return matches.length ? matches[0].index : -1;
  },

  [Symbol.split](str, limit) {
    const s = String(str);
    const lim = limit === undefined ? 2 ** 32 - 1 : limit >>> 0;
    if (lim === 0) return [];
    const matches = findPalindromes(s);
    const out = [];
    let last = 0;
    for (const m of matches) {
      out.push(s.slice(last, m.index));
      if (out.length === lim) return out;
      last = m.index + m.word.length;
    }
    out.push(s.slice(last));
    return out;
  },
};