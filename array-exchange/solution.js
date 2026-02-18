function exchangeArrays(a, b) {
  const aCopy = a.slice().reverse();
  const bCopy = b.slice().reverse();

  a.length = 0;
  b.length = 0;

  for (let i = 0; i < bCopy.length; i++) a.push(bCopy[i]);
  for (let i = 0; i < aCopy.length; i++) b.push(aCopy[i]);
}