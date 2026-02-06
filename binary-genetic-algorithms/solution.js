var GeneticAlgorithm = function () {};

GeneticAlgorithm.prototype.generate = function(length) {
  let chromosome = '';
  for (let i = 0; i < length; i++) {
    chromosome += Math.random() < 0.5 ? '0' : '1';
  }
  return chromosome;
};

GeneticAlgorithm.prototype.select = function(population, fitnesses) {
  const totalFitness = fitnesses.reduce((a, b) => a + b, 0);

  const pick = () => {
    let r = Math.random() * totalFitness;
    for (let i = 0; i < population.length; i++) {
      r -= fitnesses[i];
      if (r <= 0) return population[i];
    }
    return population[population.length - 1];
  };

  return [pick(), pick()];
};

GeneticAlgorithm.prototype.mutate = function(chromosome, p) {
  let result = '';
  for (let i = 0; i < chromosome.length; i++) {
    if (Math.random() < p) {
      result += chromosome[i] === '0' ? '1' : '0';
    } else {
      result += chromosome[i];
    }
  }
  return result;
};

GeneticAlgorithm.prototype.crossover = function(chromosome1, chromosome2) {
  const point = Math.floor(Math.random() * chromosome1.length);
  return [
    chromosome1.slice(0, point) + chromosome2.slice(point),
    chromosome2.slice(0, point) + chromosome1.slice(point)
  ];
};

GeneticAlgorithm.prototype.run = function(fitness, length, p_c, p_m, iterations) {
  if (iterations === undefined) iterations = 100;
  const popSize = 200;

  let population = [];
  for (let i = 0; i < popSize; i++) {
    population.push(this.generate(length));
  }

  for (let iter = 0; iter < iterations; iter++) {
    const fitnesses = population.map(c => fitness(c));
    const newPopulation = [];

    for (let i = 0; i < popSize / 2; i++) {
      let [c1, c2] = this.select(population, fitnesses);

      if (Math.random() < p_c) {
        [c1, c2] = this.crossover(c1, c2);
      }

      c1 = this.mutate(c1, p_m);
      c2 = this.mutate(c2, p_m);

      newPopulation.push(c1, c2);
    }

    population = newPopulation;
  }

  const fitnesses = population.map(c => fitness(c));
  let bestIdx = 0;
  for (let i = 1; i < population.length; i++) {
    if (fitnesses[i] > fitnesses[bestIdx]) bestIdx = i;
  }

  return population[bestIdx];
};