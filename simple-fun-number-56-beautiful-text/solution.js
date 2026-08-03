function isBeautifulText(s, l, r) {
  for (let width = l; width <= r; width++) {
    if (canFormatWithWidth(s, width)) {
      return true;
    }
  }
  return false;
}

function canFormatWithWidth(s, width) {
  const words = s.split(' ');
  
  // check if any word is longer than width (impossible to fit)
  for (const word of words) {
    if (word.length > width) {
      return false;
    }
  }
  
  let currentLine = '';
  
  for (const word of words) {
    if (currentLine === '') {
      currentLine = word;
    } else {
      const withSpace = currentLine + ' ' + word;
      if (withSpace.length <= width) {
        currentLine = withSpace;
      } else {
        // current line is full, check if it matches the width
        if (currentLine.length !== width) {
          return false;
        }
        currentLine = word;
      }
    }
  }
  
  // check the last line
  return currentLine.length === width;
}