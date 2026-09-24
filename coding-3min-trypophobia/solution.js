function trypophobia(arr) {
  const freq = {};
  arr.forEach(sub => sub.forEach(num => freq[num] = (freq[num] || 0) + 1));
  
  const maxFreq = Math.max(...Object.values(freq));
  const toRemove = new Set();
  
  for (let [num, count] of Object.entries(freq)) {
    if (count === maxFreq) toRemove.add(+num);
  }
  
  return arr.map(sub => sub.filter(num => !toRemove.has(num)));
}