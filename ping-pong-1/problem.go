/*
The folly of Mr Pong

When Mr Ping hits the ping pong ball, the ping pong ball goes `ping`.
When Mr Pong hits the ping pong ball, the ping pong ball goes `pong`.

Sometimes the ping pong ball hits the net, or bounces off a wall, or it ricochets
off Mrs Pong's rather expensive porcelain collection, a light fitting, or various
pieces of lounge furniture, before finally coming to rest on the floor.

## Kata Task

Who scored the most points?

## Input

The sounds of one or more rallies.

## Output

The name of the winning player: `ping` or `pong`

## Notes

- To score a point the same player must both serve and win the rally.
- A bad noise is anything other than `ping` or `pong`.
- A bad noise means whoever hit the ball last hit a bad shot.
- All noises are 4 lowercase characters, delimited by `-`
- If scores are even, then the winner is the player who did NOT hit the final bad shot.
- The final rally always ends with the ball on the floor.
*/
package main

import "strings"

func PingPong(sounds string) string {
	// stub
	return ""
}