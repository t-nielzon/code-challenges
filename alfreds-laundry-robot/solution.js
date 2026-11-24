// solution.js
function alfredLaundryRobot(commands) {
  let x = 0, y = 0;
  
  for (let cmd of commands) {
    if (cmd === 'n') y++;
    else if (cmd === 's') y--;
    else if (cmd === 'e') x++;
    else if (cmd === 'w') x--;
  }
  
  // Launderette 1 is at (3, 2), Launderette 2 is at (-4, 3)
  return (x === 3 && y === 2) || (x === -4 && y === 3);
}