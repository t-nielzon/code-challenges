/*
 * Teknonymy
 *
 * In some societies, once you have at least one child, people give you a teknonym -
 * a way to refer to someone according to some of its descendants, like the
 * 'mother of Joe', the 'great-grandfather of Jane', etc.
 *
 * Rules:
 * - If X doesn't have any direct descendant: they don't have a teknonym.
 * - If X has direct descendants: the teknonym is built according to the sex of X
 *   and the name of the elder among its direct descendants of the most distant
 *   generation from X.
 *   1) Find the most distant generation
 *   2) Find the elder of that generation (earliest dateOfBirth)
 *   3) Build the teknonym accordingly
 *
 * Teknonym format (all lowercase):
 *   'father of Y', 'grandfather of Y', 'great-grandfather of Y',
 *   'great-great-grandfather of Y', etc.
 *   'mother of Y', 'grandmother of Y', 'great-grandmother of Y', etc.
 *
 * Input: A tree where each node has:
 *   dateOfBirth : Date
 *   name : string
 *   teknonym : string  // mutate this field
 *   sex : string       // 'm' or 'f'
 *   children : Array   // possibly empty
 *
 * Output: Nothing. Mutate the teknonym field in place.
 */

function teknpionymize(tree) {
}