function rotateRight(str, n) {
  if (str.length === 0) return str;
  n = n % str.length;
  return str.slice(-n) + str.slice(0, -n);
}

function rotateLeft(str, n) {
  if (str.length === 0) return str;
  n = n % str.length;
  return str.slice(n) + str.slice(0, n);
}

const IterativeRotationCipher = {
  encode(n, str) {
    for (let i = 0; i < n; i++) {
      // Step 1: remove spaces
      let noSpace = '';
      for (let j = 0; j < str.length; j++) {
        if (str[j] !== ' ') {
          noSpace += str[j];
        }
      }
      
      // Step 2: rotate right
      noSpace = rotateRight(noSpace, n);
      
      // Step 3: put spaces back
      let withSpace = '';
      let noSpaceIdx = 0;
      for (let j = 0; j < str.length; j++) {
        if (str[j] === ' ') {
          withSpace += ' ';
        } else {
          withSpace += noSpace[noSpaceIdx++];
        }
      }
      
      // Step 4: rotate each substring right
      const parts = withSpace.split(/( +)/);
      str = parts.map(part => {
        if (part.match(/^ +$/)) return part;
        return rotateRight(part, n);
      }).join('');
    }
    
    return n + ' ' + str;
  },
  
  decode(str) {
    const spaceIdx = str.indexOf(' ');
    const n = parseInt(str.slice(0, spaceIdx));
    let result = str.slice(spaceIdx + 1);
    
    for (let i = 0; i < n; i++) {
      // Reverse step 4: rotate each substring left
      const parts = result.split(/( +)/);
      result = parts.map(part => {
        if (part.match(/^ +$/)) return part;
        return rotateLeft(part, n);
      }).join('');
      
      // Reverse step 3: remember space positions and remove spaces
      const spacePositions = [];
      let noSpace = '';
      for (let j = 0; j < result.length; j++) {
        if (result[j] === ' ') {
          spacePositions.push(j);
        } else {
          noSpace += result[j];
        }
      }
      
      // Reverse step 2: rotate left
      noSpace = rotateLeft(noSpace, n);
      
      // Reverse step 1: put spaces back
      const newResult = [];
      let noSpaceIdx = 0;
      for (let j = 0; j < result.length; j++) {
        if (spacePositions.includes(j)) {
          newResult.push(' ');
        } else {
          newResult.push(noSpace[noSpaceIdx++]);
        }
      }
      result = newResult.join('');
    }
    
    return result;
  }
};