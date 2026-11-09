function oddest(arr) {
  function oddness_level(n, seen = {}) {
    if (n % 2 === 0) return 0;
    
    if (seen[n]) return Infinity;
    
    seen[n] = true;
    const m = (n - 1) / 2;
    const m_oddness = oddness_level(m, seen);
    
    if (m_oddness === Infinity) return Infinity;
    return 1 + m_oddness;
  }
  
  if (arr.length === 0) return null;
  
  const oddness_values = arr.map(n => oddness_level(n));
  const max_oddness = Math.max(...oddness_values);
  
  const oddest_nums = arr.filter((n, i) => oddness_values[i] === max_oddness);
  
  if (oddest_nums.length === 1) {
    return oddest_nums[0];
  } else {
    return null;
  }
}