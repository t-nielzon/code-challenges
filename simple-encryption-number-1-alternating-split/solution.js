function encrypt(s, n) {
  if (!s || n <= 0) return s;
  
  let result = s;
  for (let i = 0; i < n; i++) {
    let odd = '';
    let even = '';
    for (let j = 0; j < result.length; j++) {
      if (j % 2 === 0) {
        even += result[j];
      } else {
        odd += result[j];
      }
    }
    result = odd + even;
  }
  return result;
}

function decrypt(s, n) {
  if (!s || n <= 0) return s;
  
  let result = s;
  for (let i = 0; i < n; i++) {
    const len = result.length;
    const oddCount = Math.floor(len / 2);
    const evenCount = Math.ceil(len / 2);
    
    const odd = result.substring(0, oddCount);
    const even = result.substring(oddCount);
    
    let decrypted = '';
    for (let j = 0; j < evenCount; j++) {
      decrypted += even[j];
      if (j < odd.length) {
        decrypted += odd[j];
      }
    }
    
    result = decrypted;
  }
  return result;
}