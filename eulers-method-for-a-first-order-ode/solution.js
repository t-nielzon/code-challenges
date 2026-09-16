function ex_euler(n) {
  const h = 1 / n;
  let x = 0;
  let y = 1;
  
  const f = (x, y) => 2 - Math.exp(-4 * x) - 2 * y;
  const z = (x) => 1 + 0.5 * Math.exp(-4 * x) - 0.5 * Math.exp(-2 * x);
  
  let sumErrors = 0;
  
  // error at starting point
  sumErrors += Math.abs(y - z(x)) / z(x);
  
  // perform n steps
  for (let i = 0; i < n; i++) {
    y = y + f(x, y) * h;
    x = x + h;
    sumErrors += Math.abs(y - z(x)) / z(x);
  }
  
  // mean of errors across n+1 points
  const meanError = sumErrors / (n + 1);
  
  // truncate to 6 decimal places
  return Math.floor(meanError * 1000000) / 1000000;
}