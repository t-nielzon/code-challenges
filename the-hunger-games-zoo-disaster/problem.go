package main

// Story
// A freak power outage at the zoo has caused all of the electric cage doors to malfunction and swing open...
// All the animals are out and some of them are eating each other!
//
// Here is a list of zoo animals, and what they can eat:
// - antelope eats grass
// - big-fish eats little-fish
// - bug eats leaves
// - bear eats big-fish, bug, chicken, cow, leaves, sheep
// - chicken eats bug
// - cow eats grass
// - fox eats chicken, sheep
// - giraffe eats leaves
// - lion eats antelope, cow
// - panda eats leaves
// - sheep eats grass
//
// Task: Figure out who eats whom until no more eating is possible.
//
// Rules:
// - Animals can only eat things beside them (adjacent)
// - Animals always eat to their LEFT before eating to their RIGHT
// - Always the LEFTMOST animal capable of eating will eat before any others
// - Any other things you may find at the zoo (which are not listed above) do not eat anything and are not edible
//
// Input: A comma-separated string representing all the things at the zoo
// Output: A list of strings where:
//   - The first element is the initial zoo (same as INPUT)
//   - The last element is a comma-separated string of what the zoo looks like when all eating is finished
//   - All other elements (2nd to last-1) are of the form "X eats Y" describing what happened

func ZooDis(zoo string) []string {
	// TODO: implement
	return nil
}