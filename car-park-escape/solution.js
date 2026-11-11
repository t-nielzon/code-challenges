function parkingEscape(carpark) {
  const moves = [];
  
  // find starting position (value 2)
  let currentLevel = -1;
  let currentPos = -1;
  
  for (let i = 0; i < carpark.length; i++) {
    for (let j = 0; j < carpark[i].length; j++) {
      if (carpark[i][j] === 2) {
        currentLevel = i;
        currentPos = j;
        break;
      }
    }
    if (currentLevel !== -1) break;
  }
  
  // move down levels until reaching ground floor
  while (currentLevel > 0) {
    // find staircase on current floor
    let staircasePos = -1;
    for (let j = 0; j < carpark[currentLevel].length; j++) {
      if (carpark[currentLevel][j] === 1) {
        staircasePos = j;
        break;
      }
    }
    
    // move horizontally to staircase
    if (staircasePos < currentPos) {
      const distance = currentPos - staircasePos;
      moves.push(`L${distance}`);
      currentPos = staircasePos;
    } else if (staircasePos > currentPos) {
      const distance = staircasePos - currentPos;
      moves.push(`R${distance}`);
      currentPos = staircasePos;
    }
    
    // go down one level
    moves.push("D1");
    currentLevel++;
  }
  
  // move to exit (far right of ground floor)
  const exitPos = carpark[currentLevel].length - 1;
  if (exitPos > currentPos) {
    const distance = exitPos - currentPos;
    moves.push(`R${distance}`);
  } else if (exitPos < currentPos) {
    const distance = currentPos - exitPos;
    moves.push(`L${distance}`);
  }
  
  return moves;
}