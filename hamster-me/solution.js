function hamsterMe(code, message) {
  // build a mapping: letter -> { base, distance }
  const mapping = {};

  // unique code letters in order of first appearance
  const seen = new Set();
  const codeLetters = [];
  for (const ch of code) {
    if (!seen.has(ch)) {
      seen.add(ch);
      codeLetters.push(ch);
    }
  }

  // assign distance 1 to code letters themselves
  for (const ch of codeLetters) {
    mapping[ch] = { base: ch, dist: 1 };
  }

  // for each code letter, walk forward through the alphabet assigning increasing distances
  for (const ch of codeLetters) {
    let charCode = ch.charCodeAt(0);
    let dist = 2;
    while (true) {
      charCode++;
      if (charCode > 122) charCode = 97; // wrap z -> a
      const next = String.fromCharCode(charCode);
      if (next === ch) break; // wrapped all the way around
      if (!mapping[next] || mapping[next].dist > dist) {
        mapping[next] = { base: ch, dist };
      }
      dist++;
    }
  }

  // encode message
  let result = '';
  for (const ch of message) {
    const m = mapping[ch];
    result += m.base + m.dist;
  }
  return result;
}