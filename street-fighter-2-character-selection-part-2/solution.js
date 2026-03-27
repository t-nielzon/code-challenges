function superStreetFighterIITurbo(fighters, position, moves) {
  const rows = fighters.length;
  const cols = fighters[0].length;
  let [y, x] = position;
  const result = [];

  for (const move of moves) {
    switch (move) {
      case 'up': {
        let ny = y - 1;
        // move up but stop if out of bounds or empty
        while (ny >= 0 && fighters[ny][x] === '') ny--;
        if (ny >= 0) y = ny;
        break;
      }
      case 'down': {
        let ny = y + 1;
        while (ny < rows && fighters[ny][x] === '') ny++;
        if (ny < rows) y = ny;
        break;
      }
      case 'left': {
        let nx = (x - 1 + cols) % cols;
        while (fighters[y][nx] === '') nx = (nx - 1 + cols) % cols;
        x = nx;
        break;
      }
      case 'right': {
        let nx = (x + 1) % cols;
        while (fighters[y][nx] === '') nx = (nx + 1) % cols;
        x = nx;
        break;
      }
    }
    result.push(fighters[y][x]);
  }

  return result;
}