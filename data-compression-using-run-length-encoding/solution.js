function encode(str) {
  if (str.length === 0) return '';
  
  let result = '';
  let count = 1;
  
  for (let i = 0; i < str.length; i++) {
    if (i + 1 >= str.length || str[i] !== str[i + 1]) {
      result += count + str[i];
      count = 1;
    } else {
      count++;
    }
  }
  
  return result;
}

function decode(str) {
  let result = '';
  let i = 0;
  
  while (i < str.length) {
    let num = '';
    while (i < str.length && /\d/.test(str[i])) {
      num += str[i];
      i++;
    }
    if (i < str.length) {
      result += str[i].repeat(parseInt(num));
      i++;
    }
  }
  
  return result;
}