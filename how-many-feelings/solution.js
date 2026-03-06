function countFeelings(string, array) {
  const charCount = {};
  for (const ch of string) {
    charCount[ch] = (charCount[ch] || 0) + 1;
  }

  let count = 0;
  for (const feeling of array) {
    const needed = {};
    for (const ch of feeling) {
      needed[ch] = (needed[ch] || 0) + 1;
    }
    if (Object.keys(needed).every(ch => (charCount[ch] || 0) >= needed[ch])) {
      count++;
    }
  }

  return `${count} feeling${count === 1 ? '' : 's'}.`;
}