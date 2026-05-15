/*
 * Teknonymy
 * Difficulty: 5 kyu
 *
 * In some societies, you receive a name around your birth and people use it to refer to you
 * while you don't have any children. Once you have at least one child, people give you a teknonym.
 *
 * Rules:
 * - If X doesn't have any direct descendant: they don't have a teknonym.
 * - If X has some direct descendants: the teknonym is built according to the sex of X and the
 *   name of the elder among its direct descendants of the most distant generation from X.
 *
 * Possible teknonyms (all lower case):
 *   'father of Y', 'grandfather of Y', 'great-grandfather of Y', 'great-great-grandfather of Y', ...
 *   'mother of Y', 'grandmother of Y', 'great-grandmother of Y', 'great-great-grandmother of Y', ...
 *
 * Input: a family tree node with { name, dateOfBirth, sex, teknonym, children }.
 * Output: nothing. Mutate the teknonym field where relevant.
 */
function teknonymy(person) {
  // your code here
}