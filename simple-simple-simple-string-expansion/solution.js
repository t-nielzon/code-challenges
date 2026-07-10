function stringExpansion(str) {
  if (str === "") return "";
  
  let result = "";
  let multiplier = null;
  
  for (let char of str) {
    if (/\d/.test(char)) {
      multiplier = parseInt(char);
    } else {
      if (multiplier !== null) {
        result += char.repeat(multiplier);
      } else {
        result += char;
      }
    }
  }
  
  return result;
}