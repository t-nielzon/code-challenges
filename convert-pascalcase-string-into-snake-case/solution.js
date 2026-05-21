function toUnderscore(string) {
  return String(string).replace(/([A-Z])/g, (m, p1, offset) =>
    (offset > 0 ? '_' : '') + p1.toLowerCase()
  );
}