function harvesterRescue(data) {
  const harvester = data.harvester;
  const [wormPos, wormSpeed] = data.worm;
  const [carryallPos, carryallSpeed] = data.carryall;
  
  const distance = (p1, p2) => Math.sqrt((p1[0] - p2[0]) ** 2 + (p1[1] - p2[1]) ** 2);
  
  const wormDist = distance(wormPos, harvester);
  const wormTime = wormDist / wormSpeed;
  
  const carryallDist = distance(carryallPos, harvester);
  const carryallTime = carryallDist / carryallSpeed + 1;
  
  return carryallTime <= wormTime
    ? 'The spice must flow! Rescue the harvester!'
    : 'Damn the spice! I\'ll rescue the miners!';
}