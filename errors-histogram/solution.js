function hist(s) {
  const errors = ['u', 'w', 'x', 'z'];
  const counts = {};
  for (const ch of s) {
    if (errors.includes(ch)) {
      counts[ch] = (counts[ch] || 0) + 1;
    }
  }
  const lines = [];
  for (const letter of errors) {
    if (counts[letter]) {
      const n = counts[letter];
      const letterField = letter.padEnd(2, ' ');
      const numField = String(n).padEnd(6, ' ');
      lines.push(`${letterField} ${numField}${'*'.repeat(n)}`);
    }
  }
  return lines.join('\r');
}