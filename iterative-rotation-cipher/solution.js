const IterativeRotationCipher = {
  encode: function(n, str) {
    let current = str;
    for (let i = 0; i < n; i++) {
      current = this.performIteration(current, n);
    }
    return n + ' ' + current;
  },
  
  decode: function(str) {
    const spaceIdx = str.indexOf(' ');
    const n = parseInt(str.substring(0, spaceIdx));
    const encoded = str.substring(spaceIdx + 1);
    let current = encoded;
    for (let i = 0; i < n; i++) {
      current = this.performIterationReverse(current, n);
    }
    return current;
  },
  
  performIteration: function(str, n) {
    const spacePositions = new Set();
    for (let i = 0; i < str.length; i++) {
      if (str[i] === ' ') {
        spacePositions.add(i);
      }
    }
    
    let noSpaces = '';
    for (let i = 0; i < str.length; i++) {
      if (str[i] !== ' ') {
        noSpaces += str[i];
      }
    }
    
    const rotated = this.rotateRight(noSpaces, n);
    
    let withSpaces = '';
    let rotatedIdx = 0;
    for (let i = 0; i < str.length; i++) {
      if (spacePositions.has(i)) {
        withSpaces += ' ';
      } else {
        withSpaces += rotated[rotatedIdx++];
      }
    }
    
    const substrings = withSpaces.split(/( +)/);
    let result = '';
    for (let i = 0; i < substrings.length; i++) {
      if (substrings[i].match(/^ +$/)) {
        result += substrings[i];
      } else if (substrings[i] !== '') {
        result += this.rotateRight(substrings[i], n);
      }
    }
    
    return result;
  },
  
  performIterationReverse: function(str, n) {
    const substrings = str.split(/( +)/);
    let afterSubstringRotate = '';
    for (let i = 0; i < substrings.length; i++) {
      if (substrings[i].match(/^ +$/)) {
        afterSubstringRotate += substrings[i];
      } else if (substrings[i] !== '') {
        afterSubstringRotate += this.rotateLeft(substrings[i], n);
      }
    }
    
    const spacePositions = new Set();
    for (let i = 0; i < afterSubstringRotate.length; i++) {
      if (afterSubstringRotate[i] === ' ') {
        spacePositions.add(i);
      }
    }
    
    let noSpaces = '';
    for (let i = 0; i < afterSubstringRotate.length; i++) {
      if (afterSubstringRotate[i] !== ' ') {
        noSpaces += afterSubstringRotate[i];
      }
    }
    
    const rotatedBack = this.rotateLeft(noSpaces, n);
    
    let result = '';
    let rotatedIdx = 0;
    for (let i = 0; i < afterSubstringRotate.length; i++) {
      if (spacePositions.has(i)) {
        result += ' ';
      } else {
        result += rotatedBack[rotatedIdx++];
      }
    }
    
    return result;
  },
  
  rotateRight: function(str, n) {
    if (str.length === 0) return str;
    const rotateBy = n % str.length;
    if (rotateBy === 0) return str;
    return str.slice(-rotateBy) + str.slice(0, -rotateBy);
  },
  
  rotateLeft: function(str, n) {
    if (str.length === 0) return str;
    const rotateBy = n % str.length;
    if (rotateBy === 0) return str;
    return str.slice(rotateBy) + str.slice(0, rotateBy);
  }
};