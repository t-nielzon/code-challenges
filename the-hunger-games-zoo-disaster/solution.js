function zooDisaster(zoo) {
  const menu = {
    'antelope': ['grass'],
    'big-fish': ['little-fish'],
    'bug': ['leaves'],
    'bear': ['big-fish', 'bug', 'chicken', 'cow', 'leaves', 'sheep'],
    'chicken': ['bug'],
    'cow': ['grass'],
    'fox': ['chicken', 'sheep'],
    'giraffe': ['leaves'],
    'lion': ['antelope', 'cow'],
    'panda': ['leaves'],
    'sheep': ['grass']
  };

  const result = [zoo];
  let things = zoo.split(',');

  while (true) {
    let ate = false;
    for (let i = 0; i < things.length; i++) {
      const animal = things[i];
      if (!menu[animal]) continue;
      const diet = menu[animal];
      if (i > 0 && diet.includes(things[i - 1])) {
        result.push(`${animal} eats ${things[i - 1]}`);
        things.splice(i - 1, 1);
        ate = true;
        break;
      }
      if (i < things.length - 1 && diet.includes(things[i + 1])) {
        result.push(`${animal} eats ${things[i + 1]}`);
        things.splice(i + 1, 1);
        ate = true;
        break;
      }
    }
    if (!ate) break;
  }

  result.push(things.join(','));
  return result;
}