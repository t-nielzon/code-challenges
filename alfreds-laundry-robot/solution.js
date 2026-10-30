function robotPathFinding(commands) {
  let x = 0, y = 0;
  
  for (let cmd of commands) {
    if (cmd === 'n') y++;
    else if (cmd === 's') y--;
    else if (cmd === 'e') x++;
    else if (cmd === 'w') x--;
  }
  
  return (x === 3 && y === 2) || (x === -4 && y === 3);
}