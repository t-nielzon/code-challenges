function translate(str) {
  const vowelMap = {
    b:'a',c:'a',d:'a',
    f:'e',g:'e',h:'e',
    j:'i',k:'i',l:'i',m:'i',n:'i',
    p:'o',q:'o',r:'o',s:'o',t:'o',
    v:'u',w:'u',x:'u',y:'u',z:'u'
  };

  return str.replace(/./g, ch => {
    const lower = ch.toLowerCase();
    if (vowelMap[lower]) return ch + vowelMap[lower];
    return ch;
  });
}