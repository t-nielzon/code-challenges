function pedps(a, b) {
  const result = [];
  
  let start = Math.ceil(Math.sqrt(a));
  
  for (let i = start; i * i <= b; i++) {
    const square = i * i;
    if (String(square).split('').every(digit => digit % 2 === 0)) {
      result.push(square);
    }
  }
  
  return result;
}