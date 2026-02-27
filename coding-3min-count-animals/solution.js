const names = ["dog","cat","bat","cock","cow","pig","fox",
               "ant","bird","lion","wolf","deer","bear","frog",
               "hen","mole","duck","goat"];

function sc(chars) {
  const freq = {};
  for (const c of chars) freq[c] = (freq[c] || 0) + 1;

  const animalFreqs = names.map(name => {
    const f = {};
    for (const c of name) f[c] = (f[c] || 0) + 1;
    return f;
  });

  function canForm(f, af) {
    for (const c in af) if ((f[c] || 0) < af[c]) return false;
    return true;
  }

  let max = 0;

  function search(f, count, start) {
    let remaining = 0;
    for (const c in f) remaining += f[c];
    if (count + Math.floor(remaining / 3) <= max) return;

    if (count > max) max = count;

    for (let i = start; i < names.length; i++) {
      if (canForm(f, animalFreqs[i])) {
        for (const c in animalFreqs[i]) f[c] -= animalFreqs[i][c];
        search(f, count + 1, i);
        for (const c in animalFreqs[i]) f[c] += animalFreqs[i][c];
      }
    }
  }

  search(freq, 0, 0);
  return max;
}