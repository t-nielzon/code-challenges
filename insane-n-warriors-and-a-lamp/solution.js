function shortestTime(times) {
  if (times.length === 1) return times[0];
  if (times.length === 2) return Math.max(times[0], times[1]);
  
  const sorted = times.slice().sort((a, b) => a - b);
  
  let total = 0;
  let left = sorted.length;
  
  while (left > 3) {
    const fastest = sorted[0];
    const secondFastest = sorted[1];
    const slowest = sorted[left - 1];
    const secondSlowest = sorted[left - 2];
    
    // Strategy 1: fastest person shuttles both slowest people across
    const strategyFastestShuttles = slowest + secondSlowest + 2 * fastest;
    
    // Strategy 2: two fastest people go, one returns; two slowest cross together; second fastest returns
    const strategyTwoFastestHelp = slowest + fastest + 2 * secondFastest;
    
    total += Math.min(strategyFastestShuttles, strategyTwoFastestHelp);
    left -= 2;
  }
  
  if (left === 3) {
    // with three people: fastest takes second, returns, then takes third
    total += sorted[0] + sorted[1] + sorted[2];
  } else if (left === 2) {
    // with two people: both cross together
    total += sorted[1];
  } else if (left === 1) {
    // with one person: they cross alone
    total += sorted[0];
  }
  
  return total;
}