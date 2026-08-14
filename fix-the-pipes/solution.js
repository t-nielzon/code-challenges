function fixPipes(map, start, end) {
  const height = map.length;
  const width = map[0].length;
  
  const path = [];
  
  // Determine snake pattern direction
  let goDownFirst;
  if (start < end) {
    goDownFirst = true;
  } else if (start > end) {
    goDownFirst = false;
  } else {
    goDownFirst = start < height / 2;
  }
  
  // Build the path through the grid using a snake pattern
  for (let col = 0; col < width; col++) {
    if (goDownFirst) {
      if (col % 2 === 0) {
        for (let row = 0; row < height; row++) {
          path.push([row, col]);
        }
      } else {
        for (let row = height - 1; row >= 0; row--) {
          path.push([row, col]);
        }
      }
    } else {
      if (col % 2 === 0) {
        for (let row = height - 1; row >= 0; row--) {
          path.push([row, col]);
        }
      } else {
        for (let row = 0; row < height; row++) {
          path.push([row, col]);
        }
      }
    }
  }
  
  const result = map.map(r => r.split(''));
  
  // Fill each cell with appropriate pipe
  for (let i = 0; i < path.length; i++) {
    const [row, col] = path[i];
    let fromDir, toDir;
    
    if (i === 0) {
      fromDir = 'LEFT';
    } else {
      const [prevRow] = path[i - 1];
      if (row < prevRow) fromDir = 'DOWN';
      else if (row > prevRow) fromDir = 'UP';
      else fromDir = 'LEFT';
    }
    
    if (i === path.length - 1) {
      toDir = 'RIGHT';
    } else {
      const [nextRow] = path[i + 1];
      if (row < nextRow) toDir = 'DOWN';
      else if (row > nextRow) toDir = 'UP';
      else toDir = 'RIGHT';
    }
    
    const pipe = getPipe(fromDir, toDir);
    result[row][col] = pipe;
  }
  
  return result.map(r => r.join(''));
}

function getPipe(from, to) {
  const pipes = {
    'LEFT_RIGHT': '━',
    'LEFT_UP': '┛',
    'LEFT_DOWN': '┓',
    'UP_RIGHT': '┏',
    'DOWN_RIGHT': '┗',
    'UP_DOWN': '┃',
    'DOWN_UP': '┃',
  };
  return pipes[from + '_' + to] || '?';
}