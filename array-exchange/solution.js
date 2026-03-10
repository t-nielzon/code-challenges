function exchangeArrays(a, b) {
  const aCopy = a.slice().reverse();
  const bCopy = b.slice().reverse();
  a.length = 0;
  b.length = 0;
  a.push(...bCopy);
  b.push(...aCopy);
}