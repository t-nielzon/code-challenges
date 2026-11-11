package main

/*
Directions Reduction (5 kyu)

Once upon a time, on a way through the old wild mountainous west,
a man was given directions to go from one point to another. The directions
were "NORTH", "SOUTH", "WEST", "EAST". Clearly "NORTH" and "SOUTH" are
opposite, "WEST" and "EAST" too.

Going to one direction and coming back the opposite direction right away
is a needless effort. Since this is the wild west, with dreadful weather
and not much water, it's important to save yourself some energy, otherwise
you might die of thirst!

For example:
["NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"]
becomes ["WEST"]

["NORTH", "SOUTH", "EAST", "WEST"]
becomes []

["NORTH", "EAST", "WEST", "SOUTH", "WEST", "WEST"]
becomes ["WEST", "WEST"]

Task: Write a function DirReduc which will take an array of strings and
returns an array of strings with the needless directions removed (W<->E or
S<->N side by side).

Note: Not all paths can be made simpler. The path ["NORTH", "WEST", "SOUTH",
"EAST"] is not reducible.
*/

func DirReduc(arr []string) []string {
	// TODO: implement
	return nil
}