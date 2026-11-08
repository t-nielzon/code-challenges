function sumOfSequence(begin, end, step) {
  if (begin === end) return begin;
  
  if (step > 0 && begin > end) return 0;
  if (step < 0 && begin < end) return 0;
  
  const numTerms = Math.floor(Math.abs(end - begin) / Math.abs(step)) + 1;
  const lastTerm = begin + (numTerms - 1) * step;
  
  return numTerms * (begin + lastTerm) / 2;
}