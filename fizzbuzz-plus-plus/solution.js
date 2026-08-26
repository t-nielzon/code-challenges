function fizzbuzzPlusPlus(numbers, words) {
  // LCM of coprime numbers is their product
  const lcm = numbers.reduce((a, b) => a * b, 1);
  
  const result = [];
  
  for (let i = 1; i <= lcm; i++) {
    let output = '';
    
    for (let j = 0; j < numbers.length; j++) {
      if (i % numbers[j] === 0) {
        output += words[j];
      }
    }
    
    result.push(output || i);
  }
  
  return result;
}