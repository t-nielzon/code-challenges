/*
Max wants to improve his chess skills. To achieve his goal he decided to do exercises.
Help Max at his (K)nightmares by solving this Kata.

## Task

Your task will be to return the minimum `N` (amount of moves) that are needed
on a `8` x `8` board starting with a white pawn from the initial position `start`
("a2" to "h7", 1st and 8th rank excluded) to target position `target` ("a1" to "h8").
If the `start` and `target` are equal then return `0`.

## Notes

- There are no pieces on the board except for the pawn on the starting position `start`
- A pawn can only move one square up.
- A pawn on the 2nd rank can NOT move 2 squares once. So this special chess rule will
  not take effect in this Kata.
- If the pawn reaches the 8th rank it will be promoted to a Queen or a Knight.
  It is not required to promote if the pawn can reach the target square only by moving forward.
- The Queen can move into all directions with no limits. Her only disadvantage is that
  she can not imitate a Knights movement.
*/

function underpromote(start, target) {

}