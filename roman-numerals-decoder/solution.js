function romanDecode(roman) {
  const romanValues = {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000
  };
  
  let total = 0;
  
  for (let i = 0; i < roman.length; i++) {
    const currentValue = romanValues[roman[i]];
    const nextValue = romanValues[roman[i + 1]];
    
    // if next value is greater, subtract current (subtractive notation)
    if (nextValue > currentValue) {
      total -= currentValue;
    } else {
      total += currentValue;
    }
  }
  
  return total;
}