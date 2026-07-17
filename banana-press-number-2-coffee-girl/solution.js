function minCoffeeSwaps(tray, deliverList) {
  // rotate tray 90 degrees clockwise
  function rotate90(tray) {
    const rows = tray.length;
    const cols = tray[0].length;
    const rotated = Array(cols).fill(null).map(() => Array(rows));
    
    for (let i = 0; i < rows; i++) {
      for (let j = 0; j < cols; j++) {
        rotated[j][rows - 1 - i] = tray[i][j];
      }
    }
    return rotated;
  }
  
  // flatten tray in specified reading direction
  function flattenTray(tray, readingDirection) {
    const flat = [];
    for (let i = 0; i < tray.length; i++) {
      if (readingDirection === 'ltr') {
        flat.push(...tray[i]);
      } else { // rtl
        for (let j = tray[i].length - 1; j >= 0; j--) {
          flat.push(tray[i][j]);
        }
      }
    }
    return flat;
  }
  
  // calculate minimum swaps to match target array
  function minSwaps(arr, target) {
    const temp = [...arr];
    let swaps = 0;
    
    for (let i = 0; i < temp.length; i++) {
      if (temp[i] === target[i]) continue;
      
      // find position of target[i]
      let j = i + 1;
      while (j < temp.length && temp[j] !== target[i]) {
        j++;
      }
      
      if (j < temp.length) {
        [temp[i], temp[j]] = [temp[j], temp[i]];
        swaps++;
      }
    }
    
    return swaps;
  }
  
  let minSwapsNeeded = Infinity;
  let currentTray = tray;
  
  // try all 4 rotations and both reading directions
  for (let rotation = 0; rotation < 4; rotation++) {
    for (const direction of ['ltr', 'rtl']) {
      const flattened = flattenTray(currentTray, direction);
      const swaps = minSwaps(flattened, deliverList);
      minSwapsNeeded = Math.min(minSwapsNeeded, swaps);
    }
    currentTray = rotate90(currentTray);
  }
  
  return minSwapsNeeded;
}