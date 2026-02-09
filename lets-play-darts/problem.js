/*
 * Let's Play Darts!
 *
 * Create your own mechanical dartboard that gives back your score based on
 * the coordinates of your dart.
 *
 * Task:
 * - Use the scoring rules for a standard dartboard.
 * - The coordinates (x, y) are always relative to the center of the board (0, 0).
 *   The unit is millimeters.
 *
 * Possible scores:
 * - Outside of the board: "X"
 * - Bull's eye: "DB"
 * - Bull: "SB"
 * - A single number, example: "10"
 * - A triple number: "T10"
 * - A double number: "D10"
 *
 * Circle diameters:
 * - Bull's eye: 12.70 mm
 * - Bull: 31.8 mm
 * - Triple ring inner circle: 198 mm
 * - Triple ring outer circle: 214 mm
 * - Double ring inner circle: 324 mm
 * - Double ring outer circle: 340 mm
 */

function getDartboardScore(x, y) {
}

module.exports = getDartboardScore;