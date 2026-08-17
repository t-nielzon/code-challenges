function desertRace(camels, events) {
  const camelMap = new Map();
  camels.forEach(camel => {
    camelMap.set(camel.id, camel);
  });
  
  const positions = new Map();
  const hasOasis = new Map();
  
  camels.forEach(camel => {
    positions.set(camel.id, 0);
    hasOasis.set(camel.id, false);
  });
  
  events.forEach(event => {
    const { runnerId, type, value } = event;
    
    if (!camelMap.has(runnerId)) {
      return;
    }
    
    if (type === 'move') {
      let moveValue = value;
      if (hasOasis.get(runnerId)) {
        moveValue *= 2;
        hasOasis.set(runnerId, false);
      }
      positions.set(runnerId, positions.get(runnerId) + moveValue);
    } else if (type === 'sandstorm') {
      positions.set(runnerId, positions.get(runnerId) - 2);
    } else if (type === 'oasis') {
      hasOasis.set(runnerId, true);
    }
  });
  
  let maxPosition = -Infinity;
  positions.forEach(pos => {
    maxPosition = Math.max(maxPosition, pos);
  });
  
  const winners = [];
  camels.forEach(camel => {
    if (positions.get(camel.id) === maxPosition) {
      winners.push(camel.name);
    }
  });
  
  winners.sort();
  
  return winners;
}