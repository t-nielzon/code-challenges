function desertRace(camels, events) {
  const camelMap = {};
  const positions = {};
  const hasOasis = {};
  
  camels.forEach(camel => {
    camelMap[camel.id] = camel;
    positions[camel.id] = 0;
    hasOasis[camel.id] = false;
  });
  
  events.forEach(event => {
    if (!camelMap[event.runnerId]) return;
    
    const runnerId = event.runnerId;
    
    if (event.type === 'move') {
      let moveValue = event.value;
      if (hasOasis[runnerId]) {
        moveValue *= 2;
        hasOasis[runnerId] = false;
      }
      positions[runnerId] += moveValue;
    } else if (event.type === 'sandstorm') {
      positions[runnerId] -= 2;
    } else if (event.type === 'oasis') {
      hasOasis[runnerId] = true;
    }
  });
  
  const maxPosition = Math.max(...camels.map(camel => positions[camel.id]));
  
  return camels
    .filter(camel => positions[camel.id] === maxPosition)
    .map(camel => camel.name)
    .sort();
}