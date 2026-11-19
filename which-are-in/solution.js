// solution.js
function inArray(a1, a2) {
  return [...new Set(a1.filter(str => a2.some(s => s.includes(str))))].sort();
}