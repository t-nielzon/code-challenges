function tripleTrouble(num1, num2) {
  const str1 = String(num1);
  const str2 = String(num2);
  
  // find all digits that appear as triple (3+ consecutive) in num1
  const triplesInNum1 = new Set();
  for (let i = 0; i < str1.length - 2; i++) {
    if (str1[i] === str1[i + 1] && str1[i + 1] === str1[i + 2]) {
      triplesInNum1.add(str1[i]);
    }
  }
  
  // check if any triple digit appears as double (2+ consecutive) in num2
  for (const digit of triplesInNum1) {
    for (let i = 0; i < str2.length - 1; i++) {
      if (str2[i] === digit && str2[i + 1] === digit) {
        return 1;
      }
    }
  }
  
  return 0;
}