function iterPi(epsilon) {
  let piOver4 = 0;
  let iterations = 0;
  let sign = 1;
  let denominator = 1;
  
  while (true) {
    piOver4 += sign / denominator;
    iterations++;
    
    const piApprox = piOver4 * 4;
    
    if (Math.abs(piApprox - Math.PI) < epsilon) {
      // Round to 10 decimal places
      const roundedPi = Math.round(piApprox * 1e10) / 1e10;
      return [iterations, roundedPi];
    }
    
    sign *= -1;
    denominator += 2;
  }
}