function arrange(s) {
  if (!s) return "";

  const words = s.split(/\s+/).filter(w => w.length > 0);

  // single greedy left-to-right pass: at each position enforce the required
  // relation by swapping the consecutive pair when violated. swapping never
  // breaks an already-satisfied earlier position, so one pass yields property (P)
  // and reproduces the described "carry the moved word rightward" process.
  for (let i = 0; i < words.length - 1; i++) {
    const a = words[i].length;
    const b = words[i + 1].length;
    const needLessOrEqual = i % 2 === 0; // valleys at even indices, peaks at odd

    if ((needLessOrEqual && a > b) || (!needLessOrEqual && a < b)) {
      [words[i], words[i + 1]] = [words[i + 1], words[i]];
    }
  }

  return words
    .map((w, i) => (i % 2 === 0 ? w.toLowerCase() : w.toUpperCase()))
    .join(" ");
}