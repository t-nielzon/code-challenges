function twoSum(numbers, target) {
  const map = {};
  
  for (let i = 0; i < numbers.length; i++) {
    const complement = target - numbers[i];
    
    if (complement in map) {
      return [map[complement], i];
    }
    
    map[numbers[i]] = i;
  }
}