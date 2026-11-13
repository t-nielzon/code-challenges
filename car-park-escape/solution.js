function carParkEscape(carpark) {
  const result = [];
  let currentFloor = -1;
  let currentPos = -1;
  
  // find starting position (where 2 is located)
  for (let floor = 0; floor < carpark.length; floor++) {
    const pos = carpark[floor].indexOf(2);
    if (pos !== -1) {
      currentFloor = floor;
      currentPos = pos;
      break;
    }
  }
  
  // move down floors until reaching ground floor
  while (currentFloor < carpark.length - 1) {
    // find staircase position on current floor
    const stairPos = carpark[currentFloor].indexOf(1);
    
    // move to staircase
    if (currentPos < stairPos) {
      result.push(`R${stairPos - currentPos}`);
    } else if (currentPos > stairPos) {
      result.push(`L${currentPos - stairPos}`);
    }
    
    currentPos = stairPos;
    
    // move down one level
    result.push("D1");
    currentFloor++;
  }
  
  // move to exit on ground floor (far right)
  const exitPos = carpark[carpark.length - 1].length - 1;
  if (currentPos < exitPos) {
    result.push(`R${exitPos - currentPos}`);
  } else if (currentPos > exitPos) {
    result.push(`L${currentPos - exitPos}`);
  }
  
  return result;
}