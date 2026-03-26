/*
 * Binary Genetic Algorithms (5 kyu)
 *
 * Genetic algorithms are a useful tool for machine learning. One simple way to
 * find a solution to a problem that would typically be too difficult to brute
 * force is through algorithms such as these.
 *
 * Your task is to implement a GeneticAlgorithm class with the following methods:
 *
 * - generate(length): Generate a random chromosome (binary string) of given length.
 * - select(population, fitnesses): Select two chromosomes using roulette wheel selection.
 * - crossover(chromosome1, chromosome2): With some probability, crossover two chromosomes
 *   at a random point.
 * - mutate(chromosome, p_m): With probability p_m, flip each bit in the chromosome.
 * - run(fitness, length, p_c, p_m, iterations): Run the genetic algorithm and return
 *   the fittest chromosome after the given number of iterations (default 100).
 *
 * Crossover probability: 0.6
 * Mutation probability: 0.002
 * The algorithm must discover a random 35-digit binary string using the provided
 * fitness function.
 */

var GeneticAlgorithm = function () {};

GeneticAlgorithm.prototype.generate = function(length) {
};

GeneticAlgorithm.prototype.select = function(population, fitnesses) {
};

GeneticAlgorithm.prototype.crossover = function(chromosome1, chromosome2) {
};

GeneticAlgorithm.prototype.mutate = function(chromosome, p) {
};

GeneticAlgorithm.prototype.run = function(fitness, length, p_c, p_m, iterations) {
};