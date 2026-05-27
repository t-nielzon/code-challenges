function decypher(str) {
  const lower = { '1': 'l', '2': 'z', '3': 'e', '4': 'a', '5': 's', '6': 'b', '7': 't', '8': 'B', '9': 'g', '0': 'o' };
  const upper = { '1': 'I', '2': 'R', '3': 'E', '4': 'A', '5': 'S', '6': 'G', '7': 'T', '8': 'B', '9': 'g', '0': 'O' };

  const chars = str.split('');
  let seenLetter = false;

  // a character is part of a "word" if it is a letter or a digit that maps to a letter
  const isLetterLike = c => c !== undefined && (/[A-Za-z]/.test(c) || c in lower);

  return chars.map((c, i) => {
    if (/[A-Za-z]/.test(c)) {
      seenLetter = true;
      return c;
    }

    if (c in lower) {
      let up = false;

      if (!seenLetter) up = true;                                  // rule 1: first letter
      if (chars[i - 1] === ' ' && chars[i - 2] === '.') up = true; // rule 2: after a period

      if (c === '1') {                                             // rule 3: standalone 'i'
        const prevSpaceOrStart = i === 0 || chars[i - 1] === ' ';
        const nextNonLetter = !isLetterLike(chars[i + 1]);
        if (prevSpaceOrStart && nextNonLetter) up = true;
      }

      seenLetter = true;
      return up ? upper[c] : lower[c];
    }

    return c;
  }).join('');
}