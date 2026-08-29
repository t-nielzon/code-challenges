function sumLudic(n) {
  if (n === 0) return 0;
  
  const ludic = [1];
  let numbers = Array.from({length: 240000}, (_, i) => i + 2);
  
  while (ludic.length < n && numbers.length > 0) {
    const step = numbers[0];
    ludic.push(step);
    numbers = numbers.slice(1);
    // Remove every step-th element: keep indices where index % step !== step - 1
    numbers = numbers.filter((_, index) => index % step !== step - 1);
  }
  
  return ludic.slice(0, n).reduce((sum, x) => sum + x, 0);
}