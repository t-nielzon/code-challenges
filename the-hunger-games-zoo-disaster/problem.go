package main

/*
The Hunger Games - Zoo Disaster!

A freak power outage at the zoo has caused all of the electric cage doors to malfunction
and swing open. All the animals are out and some of them are eating each other!

It's a Zoo Disaster!

INPUT: A comma-separated string representing all the things at the zoo

TASK: Figure out who eats whom until no more eating is possible.

OUTPUT: A list of strings where:
- The first element is the initial zoo (same as INPUT)
- The last element is a comma-separated string of what the zoo looks like when all eating has finished
- All other elements (2nd to last-1) are of the form "X eats Y" describing what happened

Notes:
- Animals can only eat things beside them
- Animals always eat to their LEFT before eating to their RIGHT
- Always the LEFTMOST animal capable of eating will eat before any others
- Any other things at the zoo (which are not listed above) do not eat anything and are not edible
*/
func ZooDis(zoo string) []string {
}