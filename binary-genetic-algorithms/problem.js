/*
 * Binary Genetic Algorithms
 *
 * Genetic algorithms are a useful tool for machine learning. One simple way to find a solution
 * to a problem that would typically be too difficult to brute force is through algorithms such as these.
 *
 * You will be given an outline of a GeneticAlgorithm class with a few methods:
 * - crossover(chromosome1, chromosome2): takes two chromosomes and returns a crossed-over pair
 * - mutate(chromosome, p): takes a chromosome and mutation probability, returns mutated chromosome
 * - generate(length): generates a random chromosome of a given length
 * - select(population, fitnesses): takes a population and fitnesses, returns two chromosomes
 *   selected via roulette wheel selection
 * - run(fitness, length, p_c, p_m, iterations): runs the genetic algorithm and returns the
 *   fittest chromosome after all iterations
 *
 * The crossover probability is 0.6 and mutation probability is 0.002.
 * Chromosome length is 35 bits. 100 iterations should suffice.
 */

var GeneticAlgorithm = function () {};

GeneticAlgorithm.prototype.generate = function(length) {
};

GeneticAlgorithm.prototype.select = function(population, fitnesses) {
};

GeneticAlgorithm.prototype.mutate = function(chromosome, p) {
};

GeneticAlgorithm.prototype.crossover = function(chromosome1, chromosome2) {
};

GeneticAlgorithm.prototype.run = function(fitness, length, p_c, p_m, iterations) {
};