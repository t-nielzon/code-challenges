function sevenAte9(arr) {
  let s = arr.join('');
  // keep rescuing: moving a 7 past a 9 can expose a new 7,8,9 sequence
  while (s.includes('789')) {
    s = s.replace('789', '897');
  }
  return s.split('').map(Number);
}