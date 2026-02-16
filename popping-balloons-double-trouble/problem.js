/*
## Task

Balloons drift from left to right toward you, and you stand on the right firing volleys of arrows. Each step, you shoot `n` arrows at the closest balloons on your side, then watch the whole line change before the next volley. After `k` such steps, return the surviving balloons as a string.

- **Input:**
  - `s`: initial string of balloons (`w,b,g,r,W,B,G,R`).
  - `n`: arrows fired per step.
  - `k`: number of steps to simulate.
- **Output:** Final balloon string (compact, no gaps).
- **Per step:**
  - Target the `n` rightmost balloons, or fewer if less than `n` balloons remain.
  - Hit each once; apply effects at end of step.
- **Notes:**
  - All balloons start with 0 hits.
  - Each balloon tracks hits independently.
  - Balloons always stay compact—no spaces form when some pop, transform or split.
- **Constraints:**
  - `1 <= n <= 8`
  - `0 <= k <= 1000`
  - `0 <= length of s <= 200`

## Regular Balloons

| Balloon          | Effect          |
|------------------|-----------------|
| White (w)        | pops            |
| Black (b)        | → `w`           |
| Green (g)        | → `ww`          |
| Red (r)          | → `bb`          |

## Fortified Balloons

| Balloon          | 1st Hit     | 2nd Hit       |
|------------------|-------------|---------------|
| White (W)        | → `w*`     | pops          |
| Black (B)        | → `b*`     | → `W`         |
| Green (G)        | → `g*`     | → `WW`        |
| Red (R)          | → `r*`     | → `BB`        |
*/

function popBalloons(s, n, k) {
}