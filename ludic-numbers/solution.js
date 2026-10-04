function sumLudic(n) {
  let ludic = [1];
  let numbers = [];
  
  // Generate enough numbers to find the first n ludic numbers
  for (let i = 2; i <= 250000; i++) {
    numbers.push(i);
  }
  
  // Generate ludic numbers using sieve
  while (ludic.length < n && numbers.length > 0) {
    const first = numbers[0];
    ludic.push(first);
    // Remove every first-th indexed element (keep indices where i % first !== 0)
    numbers = numbers.filter((_, i) => i % first !== 0);
  }
  
  // Sum the first n ludic numbers
  let sum = 0;
  for (let i = 0; i < n; i++) {
    sum += ludic[i];
  }
  return sum;
}