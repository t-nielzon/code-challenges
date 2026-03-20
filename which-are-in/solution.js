function inArray(array1, array2) {
  return [...new Set(array1.filter(s => array2.some(w => w.includes(s))))].sort();
}

module.exports = { inArray };