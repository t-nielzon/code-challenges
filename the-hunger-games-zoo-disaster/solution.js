const DIET = {
  antelope: ['grass'],
  'big-fish': ['little-fish'],
  bug: ['leaves'],
  bear: ['big-fish', 'bug', 'chicken', 'cow', 'leaves', 'sheep'],
  chicken: ['bug'],
  cow: ['grass'],
  fox: ['chicken', 'sheep'],
  giraffe: ['leaves'],
  lion: ['antelope', 'cow'],
  panda: ['leaves'],
  sheep: ['grass'],
};

function canEat(predator, prey) {
  return DIET[predator] !== undefined && DIET[predator].includes(prey);
}

function zooShuffle(zoo) {
  const animals = zoo.split(',');
  const result = [zoo];

  let ate = true;
  while (ate) {
    ate = false;

    // scan left-to-right for the leftmost animal that can eat,
    // preferring its left neighbour before its right neighbour.
    for (let i = 0; i < animals.length; i++) {
      const predator = animals[i];

      // eat to the left first
      if (i > 0 && canEat(predator, animals[i - 1])) {
        result.push(`${predator} eats ${animals[i - 1]}`);
        animals.splice(i - 1, 1);
        ate = true;
        break;
      }

      // then eat to the right
      if (i < animals.length - 1 && canEat(predator, animals[i + 1])) {
        result.push(`${predator} eats ${animals[i + 1]}`);
        animals.splice(i + 1, 1);
        ate = true;
        break;
      }
    }
  }

  result.push(animals.join(','));
  return result;
}