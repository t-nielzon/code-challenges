function sumOfSequence(begin, end, step) {
  // if begin equals end, return it regardless of step
  if (begin === end) {
    return begin;
  }
  
  // check for impossible sequences
  if ((begin > end && step > 0) || (begin < end && step < 0)) {
    return 0;
  }
  
  // calculate the number of steps to reach or approach the limit
  const numSteps = Math.floor((end - begin) / step);
  const lastValue = begin + numSteps * step;
  
  // use arithmetic sequence sum formula: n * (first + last) / 2
  const numTerms = numSteps + 1;
  return numTerms * (begin + lastValue) / 2;
}