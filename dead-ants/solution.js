function deadAnts(s) {
  // remove all complete ant sequences to get only scattered bits
  const cleaned = s.replace(/ant/g, '');
  
  // count occurrences of each letter in scattered bits
  const countA = (cleaned.match(/a/g) || []).length;
  const countN = (cleaned.match(/n/g) || []).length;
  const countT = (cleaned.match(/t/g) || []).length;
  
  // each dead ant contributes one of each letter to the scattered bits
  // the number of dead ants is the maximum count
  return Math.max(countA, countN, countT);
}