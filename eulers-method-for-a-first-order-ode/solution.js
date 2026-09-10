function ex_euler(n) {
  const h = 1 / n;
  
  let x = 0;
  let y = 1;
  
  const f = (x, y) => 2 - Math.exp(-4 * x) - 2 * y;
  const z = (x) => 1 + 0.5 * Math.exp(-4 * x) - 0.5 * Math.exp(-2 * x);
  
  let sumErrors = 0;
  const z0 = z(0);
  sumErrors += Math.abs(y - z0) / z0;
  
  for (let i = 0; i < n; i++) {
    y = y + f(x, y) * h;
    x = x + h;
    
    const zk = z(x);
    const error = Math.abs(y - zk) / zk;
    sumErrors += error;
  }
  
  const meanError = sumErrors / (n + 1);
  
  return Math.floor(meanError * 1000000) / 1000000;
}