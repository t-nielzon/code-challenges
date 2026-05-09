/*
 * The Hunger Games - Zoo Disaster!
 *
 * A freak power outage at the zoo has caused all of the electric cage doors
 * to malfunction and swing open. All the animals are out and some of them
 * are eating each other!
 *
 * Food chain:
 *   antelope eats grass
 *   big-fish eats little-fish
 *   bug eats leaves
 *   bear eats big-fish
 *   bear eats bug
 *   bear eats chicken
 *   bear eats cow
 *   bear eats leaves
 *   bear eats sheep
 *   chicken eats bug
 *   cow eats grass
 *   fox eats chicken
 *   fox eats sheep
 *   giraffe eats leaves
 *   lion eats antelope
 *   lion eats cow
 *   panda eats leaves
 *   sheep eats grass
 *
 * INPUT
 *   A comma-separated string representing all the things at the zoo.
 *
 * TASK
 *   Figure out who eats whom until no more eating is possible.
 *
 * OUTPUT
 *   A list of strings where:
 *   - First element is the initial zoo (same as input)
 *   - Last element is the final zoo as a comma-separated string
 *   - Middle elements are "X eats Y" describing each event
 *
 * Rules:
 *   - Animals can only eat things beside them.
 *   - Animals always eat to their LEFT before eating to their RIGHT.
 *   - The LEFTMOST animal capable of eating eats before any others.
 *   - Anything not in the list neither eats nor is eaten.
 */

function zooDisaster(zoo) {

}