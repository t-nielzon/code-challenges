var GeneticAlgorithm = function () {};

GeneticAlgorithm.prototype.generate = function(length) {
  let s = '';
  for (let i = 0; i < length; i++) {
    s += Math.random() < 0.5 ? '0' : '1';
  }
  return s;
};

GeneticAlgorithm.prototype.select = function(population, fitnesses) {
  const total = fitnesses.reduce((a, b) => a + b, 0);
  const pick = () => {
    let r = Math.random() * total;
    for (let i = 0; i < population.length; i++) {
      r -= fitnesses[i];
      if (r <= 0) return population[i];
    }
    return population[population.length - 1];
  };
  return [pick(), pick()];
};

GeneticAlgorithm.prototype.crossover = function(pair) {
  const point = Math.floor(Math.random() * pair[0].length);
  return [
    pair[0].slice(0, point) + pair[1].slice(point),
    pair[1].slice(0, point) + pair[0].slice(point)
  ];
};

GeneticAlgorithm.prototype.mutate = function(chromosome, p) {
  let s = '';
  for (let i = 0; i < chromosome.length; i++) {
    if (Math.random() < p) {
      s += chromosome[i] === '0' ? '1' : '0';
    } else {
      s += chromosome[i];
    }
  }
  return s;
};

GeneticAlgorithm.prototype.run = function(fitness, length, p_c, p_m, iterations) {
  iterations = iterations || 100;
  const popSize = 100;

  let population = [];
  for (let i = 0; i < popSize; i++) {
    population.push(this.generate(length));
  }

  for (let gen = 0; gen < iterations; gen++) {
    const fitnesses = population.map(c => fitness(c));
    const newPop = [];

    for (let i = 0; i < popSize / 2; i++) {
      let pair = this.select(population, fitnesses);

      if (Math.random() < p_c) {
        pair = this.crossover(pair);
      }

      newPop.push(this.mutate(pair[0], p_m));
      newPop.push(this.mutate(pair[1], p_m));
    }

    population = newPop;
  }

  let bestFit = -1;
  let best = population[0];
  for (let i = 0; i < population.length; i++) {
    const f = fitness(population[i]);
    if (f > bestFit) {
      bestFit = f;
      best = population[i];
    }
  }

  return best;
};