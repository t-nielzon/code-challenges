function detectBruteForce(logs) {
  const failureCounts = {};
  const suspiciousIPs = new Set();
  
  for (const log of logs) {
    const parts = log.split(' ');
    const ip = parts[0];
    const action = parts[1];
    
    if (action === 'LOGIN_FAIL') {
      failureCounts[ip] = (failureCounts[ip] || 0) + 1;
      if (failureCounts[ip] >= 3) {
        suspiciousIPs.add(ip);
      }
    } else if (action === 'LOGIN_SUCCESS') {
      failureCounts[ip] = 0;
    }
  }
  
  return Array.from(suspiciousIPs).sort();
}

module.exports = { detectBruteForce };