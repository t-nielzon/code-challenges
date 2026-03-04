/*
 * Teknonymy
 *
 * In some societies, once you have at least one child, people give you a teknonym —
 * a way to refer to someone according to some of its descendants, like the
 * 'mother of Joe', the 'great-grandfather of Jane', etc.
 *
 * Rules:
 * - If X doesn't have any direct descendant: no teknonym.
 * - If X has descendants: the teknonym is built according to X's sex and the name
 *   of the elder among its direct descendants of the most distant generation from X.
 * - Find the most distant generation, find the elder of that generation,
 *   build the teknonym.
 * - Teknonyms: 'father of Y', 'grandfather of Y', 'great-grandfather of Y',
 *   'great-great-grandfather of Y', etc. (and mother/grandmother/great-grandmother...)
 *
 * Input: A tree structure where each node has:
 *   dateOfBirth: Date, name: string, teknonym: string, sex: string, children: Array
 *
 * Output: Nothing. Mutate the teknonym field in place.
 */

function teknpionymize(tree) {

}