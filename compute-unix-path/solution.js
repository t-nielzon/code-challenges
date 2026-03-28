function compute(...parts) {
  const path = parts
    .map(p => p.trim())
    .filter(p => p.length > 0)
    .join('/')
    .replace(/\\/g, '/')
    .replace(/\/+/g, '/')
    .replace(/\/+$/, '');

  return path || '/';
}

module.exports = { compute };