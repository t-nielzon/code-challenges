function isAtLaunderette(commands) {
  const launderette1 = ['e', 'n', 'e', 'e', 'n'];
  const launderette2 = ['w', 'n', 'w', 'n', 'w', 'w', 'n'];
  
  function getPosition(path) {
    let x = 0, y = 0;
    for (const cmd of path) {
      if (cmd === 'n') y++;
      else if (cmd === 's') y--;
      else if (cmd === 'e') x++;
      else if (cmd === 'w') x--;
    }
    return [x, y];
  }
  
  const robotPos = getPosition(commands);
  const laundPos1 = getPosition(launderette1);
  const laundPos2 = getPosition(launderette2);
  
  return (robotPos[0] === laundPos1[0] && robotPos[1] === laundPos1[1]) ||
         (robotPos[0] === laundPos2[0] && robotPos[1] === laundPos2[1]);
}