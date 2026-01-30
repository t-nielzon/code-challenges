/*
Directions Reduction

Once upon a time, on a way through the old wild mountainous west,…

… a man was given directions to go from one point to another. The directions were
"NORTH", "SOUTH", "WEST", "EAST". Clearly "NORTH" and "SOUTH" are opposite, "WEST"
and "EAST" too.

Going to one direction and coming back the opposite direction right away is a
needless effort. Since this is the wild west, with dreadful weather and not much
water, it's important to save yourself some energy, otherwise you might die of thirst!

The task is to give to the man a simplified version of the plan by removing needless
directions (W<->E or S<->N side by side).

Examples:
- ["NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"] -> ["WEST"]
- ["NORTH", "SOUTH", "EAST", "WEST"] -> []
- ["NORTH", "EAST", "WEST", "SOUTH", "WEST", "WEST"] -> ["WEST", "WEST"]
- ["NORTH", "WEST", "SOUTH", "EAST"] -> ["NORTH", "WEST", "SOUTH", "EAST"] (not reducible)
*/
package kata

func DirReduc(arr []string) []string {
}