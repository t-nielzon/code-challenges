function sumOfSequence(begin, end, step) {
  // If begin equals end, return it
  if (begin === end) {
    return begin;
  }
  
  // Check for impossible sequences
  if ((step > 0 && begin > end) || (step < 0 && begin < end)) {
    return 0;
  }
  
  // Calculate the number of complete steps we can take
  const k = Math.floor((end - begin) / step);
  const numTerms = k + 1;
  
  // Calculate the actual last term in the sequence
  const lastTerm = begin + k * step;
  
  // Use arithmetic sequence sum formula: n * (first + last) / 2
  return (numTerms * (begin + lastTerm)) / 2;
}