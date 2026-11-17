function encrypt(text) {
  if (text === null || text === '') return text;
  
  const region = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789.,:;-?! \'()$%&"';
  
  for (let char of text) {
    if (!region.includes(char)) {
      throw new Error('Invalid character: ' + char);
    }
  }
  
  // Step 1: Case switch on every second char (positions 1, 3, 5, ...)
  let step1 = '';
  for (let i = 0; i < text.length; i++) {
    let char = text[i];
    if (i % 2 === 1 && ((char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z'))) {
      char = char === char.toUpperCase() ? char.toLowerCase() : char.toUpperCase();
    }
    step1 += char;
  }
  
  // Step 2: Index difference transformation
  let step2 = step1[0];
  for (let i = 1; i < step1.length; i++) {
    let prevIdx = region.indexOf(step1[i-1]);
    let currIdx = region.indexOf(step1[i]);
    let newIdx = ((prevIdx - currIdx) % 77 + 77) % 77;
    step2 += region[newIdx];
  }
  
  // Step 3: Mirror the first char
  let firstIdx = region.indexOf(step1[0]);
  let mirrorIdx = 76 - firstIdx;
  return region[mirrorIdx] + step2.slice(1);
}

function decrypt(encryptedText) {
  if (encryptedText === null || encryptedText === '') return encryptedText;
  
  const region = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789.,:;-?! \'()$%&"';
  
  for (let char of encryptedText) {
    if (!region.includes(char)) {
      throw new Error('Invalid character: ' + char);
    }
  }
  
  // Reverse step 3: Mirror the first char back
  let mirrorIdx = region.indexOf(encryptedText[0]);
  let firstIdx = 76 - mirrorIdx;
  let step2 = region[firstIdx] + encryptedText.slice(1);
  
  // Reverse step 2: Reverse the index difference
  let step1 = step2[0];
  for (let i = 1; i < step2.length; i++) {
    let prevIdx = region.indexOf(step1[i-1]);
    let encIdx = region.indexOf(step2[i]);
    let origIdx = ((prevIdx - encIdx) % 77 + 77) % 77;
    step1 += region[origIdx];
  }
  
  // Reverse step 1: Case switch again on positions 1, 3, 5, ...
  let text = '';
  for (let i = 0; i < step1.length; i++) {
    let char = step1[i];
    if (i % 2 === 1 && ((char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z'))) {
      char = char === char.toUpperCase() ? char.toLowerCase() : char.toUpperCase();
    }
    text += char;
  }
  
  return text;
}