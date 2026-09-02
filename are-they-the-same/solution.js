function comp(a, b){
  if (!Array.isArray(a) || !Array.isArray(b)) return false;
  if (a.length !== b.length) return false;
  
  const squaredA = a.map(x => x * x).sort((x, y) => x - y);
  const sortedB = b.sort((x, y) => x - y);
  
  return squaredA.every((val, i) => val === sortedB[i]);
}