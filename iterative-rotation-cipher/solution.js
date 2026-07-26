function rotateRight(str, amount) {
  if (str.length === 0) return str;
  amount = ((amount % str.length) + str.length) % str.length;
  if (amount === 0) return str;
  return str.slice(-amount) + str.slice(0, -amount);
}

function encodeOneIteration(str, n) {
  // step 1: remove spaces and remember positions
  const spacePositions = [];
  let noSpaces = '';
  for (let i = 0; i < str.length; i++) {
    if (str[i] === ' ') {
      spacePositions.push(i);
    } else {
      noSpaces += str[i];
    }
  }
  
  // step 2: rotate string right by n
  const rotatedString = rotateRight(noSpaces, n);
  
  // step 3: put spaces back
  let result = '';
  let rotatedIdx = 0;
  for (let i = 0; i < str.length; i++) {
    if (spacePositions.includes(i)) {
      result += ' ';
    } else {
      result += rotatedString[rotatedIdx++];
    }
  }
  
  // step 4: rotate each substring separated by spaces
  let finalResult = '';
  let currentSubstring = '';
  
  for (let i = 0; i < result.length; i++) {
    if (result[i] === ' ') {
      if (currentSubstring.length > 0) {
        finalResult += rotateRight(currentSubstring, n);
        currentSubstring = '';
      }
      finalResult += ' ';
    } else {
      currentSubstring += result[i];
    }
  }
  
  if (currentSubstring.length > 0) {
    finalResult += rotateRight(currentSubstring, n);
  }
  
  return finalResult;
}

function decodeOneIteration(str, n) {
  // reverse of step 4: rotate each substring left by n
  let afterStep4Reverse = '';
  let currentSubstring = '';
  
  for (let i = 0; i < str.length; i++) {
    if (str[i] === ' ') {
      if (currentSubstring.length > 0) {
        afterStep4Reverse += rotateRight(currentSubstring, -n);
        currentSubstring = '';
      }
      afterStep4Reverse += ' ';
    } else {
      currentSubstring += str[i];
    }
  }
  
  if (currentSubstring.length > 0) {
    afterStep4Reverse += rotateRight(currentSubstring, -n);
  }
  
  // reverse of steps 3 & 2: remove spaces, rotate left, put back
  const spacePositions = [];
  let noSpaces = '';
  for (let i = 0; i < afterStep4Reverse.length; i++) {
    if (afterStep4Reverse[i] === ' ') {
      spacePositions.push(i);
    } else {
      noSpaces += afterStep4Reverse[i];
    }
  }
  
  const rotatedString = rotateRight(noSpaces, -n);
  
  let result = '';
  let rotatedIdx = 0;
  for (let i = 0; i < afterStep4Reverse.length; i++) {
    if (spacePositions.includes(i)) {
      result += ' ';
    } else {
      result += rotatedString[rotatedIdx++];
    }
  }
  
  return result;
}

const IterativeRotationCipher = {
  encode(n, str) {
    let result = str;
    for (let i = 0; i < n; i++) {
      result = encodeOneIteration(result, n);
    }
    return n + ' ' + result;
  },
  
  decode(str) {
    const spaceIndex = str.indexOf(' ');
    const n = parseInt(str.substring(0, spaceIndex));
    let encoded = str.substring(spaceIndex + 1);
    
    for (let i = 0; i < n; i++) {
      encoded = decodeOneIteration(encoded, n);
    }
    
    return encoded;
  }
};