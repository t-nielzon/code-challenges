/*
 * Banana Press #2: Coffee Girl  (5 kyu)
 *
 * Given a `tray` (a 2d array of NxM integers) and a `deliverList`
 * (an array of distinct integers), where the integers in the tray
 * match the integers in the deliverList but are not in the same order,
 * return the minimum number of swaps required to order the coffees on
 * the tray (read row by row) so they match the deliverList.
 *
 * Before any swapping you are allowed to transform the whole tray:
 *   - rotate it (90 / 180 / 270 degrees),
 *   - mirror it,
 *   - and pick the coffees either left-to-right or right-to-left.
 *
 * Choose the transformation that minimizes the number of swaps.
 *
 * Example 1:
 *   tray = [[1,2,7],[4,5,6],[3,8,9]], deliverList = [1..9]  -> 1
 *
 * Example 2:
 *   tray = [[1,4,7],[2,5,8],[3,6,9]], deliverList = [1..9]  -> 0
 *   (rotate 90 clockwise and read right-to-left)
 */

function minCoffeeSwaps(tray, deliverList) {
  // your code here
}