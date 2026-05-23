function abbreviate(string) {
  return string.replace(/[a-zA-Z]{4,}/g, w => w[0] + (w.length - 2) + w[w.length - 1]);
}