function find(needle, haystack) {
  if (needle.length > haystack.length) {
    return -1;
  }
  
  for (let i = 0; i <= haystack.length - needle.length; i++) {
    let match = true;
    
    for (let j = 0; j < needle.length; j++) {
      if (needle[j] !== '_' && needle[j] !== haystack[i + j]) {
        match = false;
        break;
      }
    }
    
    if (match) {
      return i;
    }
  }
  
  return -1;
}