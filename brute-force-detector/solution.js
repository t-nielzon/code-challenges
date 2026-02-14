function detectBruteForce(logs) {
  const failCounts = {};
  const suspicious = new Set();

  for (const entry of logs) {
    const [ip, action] = entry.split(' ');

    if (action === 'LOGIN_FAIL') {
      failCounts[ip] = (failCounts[ip] || 0) + 1;
      if (failCounts[ip] >= 3) {
        suspicious.add(ip);
      }
    } else if (action === 'LOGIN_SUCCESS') {
      failCounts[ip] = 0;
    }
  }

  return [...suspicious].sort();
}

module.exports = { detectBruteForce };