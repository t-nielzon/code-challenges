function to_leet(text) {
  const scroll = {
    a: ['4', '@'],
    b: ['|3', '8'],
    d: ['|)', 'o|'],
    e: ['3'],
    f: ['|='],
    g: ['9', '6'],
    h: ['|-|', ']-[', '}-{', '(-)', ')-(', '#'],
    i: ['1', '!', ']['],
    j: ['_|'],
    k: ['|<', '|{'],
    l: ['|_'],
    n: ['|\\|'],
    o: ['0'],
    p: ['|2', '|D'],
    q: ['(,)'],
    r: ['|Z', '|?'],
    s: ['5', '$'],
    t: ['+', '7'],
    v: ['|/', '\\/'],
    w: ['\\^/', '//'],
    x: ['><', '}{'],
    y: ['`/'],
    z: ['(\\)']
  };
  
  const indices = {};
  let result = '';
  
  for (const char of text) {
    if (char === ' ') {
      result += ' ';
    } else if (scroll[char]) {
      if (!indices[char]) {
        indices[char] = 0;
      }
      const options = scroll[char];
      result += options[indices[char] % options.length];
      indices[char]++;
    } else {
      result += char;
    }
  }
  
  return result;
}