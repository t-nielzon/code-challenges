const names = ["dog","cat","bat","cock","cow","pig","fox",
               "ant","bird","lion","wolf","deer","bear","frog",
               "hen","mole","duck","goat"];

function sc(chars) {
  // build a frequency map from a string
  function charCount(s) {
    const m = {};
    for (const c of s) m[c] = (m[c] || 0) + 1;
    return m;
  }

  // check if animal can be formed from available chars, return remaining if so
  function tryUse(available, animal) {
    const rem = Object.assign({}, available);
    for (const c of animal) {
      if (!rem[c] || rem[c] <= 0) return null;
      rem[c]--;
    }
    return rem;
  }

  // dfs with pruning to find maximum animals
  let best = 0;

  function dfs(available, count, startIdx) {
    if (count > best) best = count;

    for (let i = startIdx; i < names.length; i++) {
      const rem = tryUse(available, names[i]);
      if (rem !== null) {
        // allow same animal again (startIdx = i, not i+1)
        dfs(rem, count + 1, i);
      }
    }
  }

  dfs(charCount(chars), 0, 0);
  return best;
}

module.exports = sc;