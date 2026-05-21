/*
Ziggurat Ride of Fortune

You and a group of explorers have found a legendary ziggurat hidden in an obscure jungle.
The outer west wall of the ziggurat consists of a series of entrance doors. When an
explorer enters through a door, they sit in a mobile cart that moves in a straight path
until it reaches a "switch" or hits a wall.

A switch re-routes the cart either left or right, depending on the state of the switch
and the cart's movement direction.

Interdimensional portals line the entire north, east, and south walls inside the ziggurat.
If a cart hits one of these walls, the occupying explorer exits through the portal before
them. If a cart ends at the west wall, the explorer exits through a door and returns back
outside.

Switch Mechanics:
If a switch is in the A state and a cart enters by moving:
  - west, they are routed north
  - east, they are routed south
  - south, they are routed east
  - north, they are routed west
If the switch were in the B state, the cart would be routed in the orthogonal direction
opposite to that for the A state.
Immediately after a cart passes through a switch, the switch changes state by rotating
90 degrees.

Input:
  - An n x n matrix representing the layout of the ziggurat interior.
  - An array of the doors (rows) entered by each explorer in sequence.

Output:
  - An array of the exit points of explorers who exit through portals, and [-1, -1] for
    those who return back outside.
*/

package solution

func RideOfFortune(artifact []string, explorers []int) [][2]int {
	// your code here
	return nil
}