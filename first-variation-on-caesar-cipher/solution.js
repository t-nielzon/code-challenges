function movingShift(s, shift) {
  let encoded = "";
  let currentShift = shift;
  
  for (let char of s) {
    if (/[a-z]/.test(char)) {
      let charCode = char.charCodeAt(0) - 'a'.charCodeAt(0);
      charCode = (charCode + currentShift) % 26;
      encoded += String.fromCharCode(charCode + 'a'.charCodeAt(0));
    } else if (/[A-Z]/.test(char)) {
      let charCode = char.charCodeAt(0) - 'A'.charCodeAt(0);
      charCode = (charCode + currentShift) % 26;
      encoded += String.fromCharCode(charCode + 'A'.charCodeAt(0));
    } else {
      encoded += char;
    }
    currentShift++;
  }
  
  let length = encoded.length;
  let baseHigh = Math.ceil(length / 5);
  let baseLow = Math.floor(length / 5);
  
  let part1 = baseHigh;
  let part2 = baseHigh;
  let part3 = baseHigh;
  let remaining = length - part1 - part2 - part3;
  
  let part4 = Math.max(baseLow, Math.min(baseHigh, remaining));
  remaining -= part4;
  
  let part5 = remaining;
  
  let parts = [];
  let index = 0;
  for (let size of [part1, part2, part3, part4, part5]) {
    parts.push(encoded.substring(index, index + size));
    index += size;
  }
  
  return parts;
}

function demovingShift(arr, shift) {
  let encoded = arr.join("");
  let decoded = "";
  let currentShift = shift;
  
  for (let char of encoded) {
    if (/[a-z]/.test(char)) {
      let charCode = char.charCodeAt(0) - 'a'.charCodeAt(0);
      let shiftMod = currentShift % 26;
      charCode = ((charCode - shiftMod) % 26 + 26) % 26;
      decoded += String.fromCharCode(charCode + 'a'.charCodeAt(0));
    } else if (/[A-Z]/.test(char)) {
      let charCode = char.charCodeAt(0) - 'A'.charCodeAt(0);
      let shiftMod = currentShift % 26;
      charCode = ((charCode - shiftMod) % 26 + 26) % 26;
      decoded += String.fromCharCode(charCode + 'A'.charCodeAt(0));
    } else {
      decoded += char;
    }
    currentShift++;
  }
  
  return decoded;
}