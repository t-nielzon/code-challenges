function temps(v0, slope, dTot) {
  const GRAVITY_ACC = 9.81 * 3.6 * 60.0;
  const DRAG = 60.0 * 0.3 / 3.6;
  const DELTA_T = 1.0 / 60.0;
  const G_THRUST = 60 * 3.6 * 3.6;
  const MASS = 80.0;
  const WATTS0 = 225.0;
  const D_WATTS = 0.5;
  
  let t = 0;
  let v = v0;
  let d = 0;
  let watts = WATTS0;
  
  while (d < dTot) {
    // Calculate acceleration components
    const slopeRad = Math.atan(slope / 100);
    const gravityComponent = -GRAVITY_ACC * Math.sin(slopeRad);
    const dragComponent = -DRAG * Math.abs(v) * Math.abs(v) / MASS;
    let thrustComponent = 0;
    
    if (v > 0 && watts > 0) {
      thrustComponent = G_THRUST * watts / (v * MASS);
    }
    
    let gamma = gravityComponent + dragComponent + thrustComponent;
    
    if (Math.abs(gamma) <= 1e-5) {
      gamma = 0;
    }
    
    // Update velocity
    v = v + gamma * DELTA_T;
    
    // Check if John gives up
    if (v - 3.0 <= 1e-2) {
      return -1;
    }
    
    // Update distance
    d = d + v * DELTA_T / 60.0;
    
    // Update power
    watts = watts - D_WATTS * DELTA_T;
    
    // Increment time
    t = t + DELTA_T;
  }
  
  return Math.round(t);
}