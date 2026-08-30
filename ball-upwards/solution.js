function maxHeight(v0) {
  const g = 9.81;
  const v0_ms = v0 / 3.6; // convert km/h to m/s
  
  let maxH = 0;
  let maxIndex = 0;
  
  for (let i = 0; ; i++) {
    const t = i * 0.1; // time in seconds
    const h = v0_ms * t - 0.5 * g * t * t;
    
    if (h < 0) break; // ball has hit the ground
    
    if (h > maxH) {
      maxH = h;
      maxIndex = i;
    }
  }
  
  return maxIndex;
}