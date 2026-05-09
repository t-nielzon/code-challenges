function inArray(a1, a2) {
  return [...new Set(a1.filter(s => a2.some(t => t.includes(s))))].sort();
}