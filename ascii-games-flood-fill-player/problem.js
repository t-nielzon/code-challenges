/*
 * ASCII Games: Flood Fill (Player)
 *
 * You will be given a rectangular map with characters 1-5 representing different colors.
 *
 * During every turn, you can choose a coordinate x,y (top left origin, x goes rightwards,
 * y goes downwards) and a color 1-5, which will flood-fill every orthogonally adjacent tile
 * of the same color as the chosen tile into the new color.
 *
 * Play back the moves and return the result of the transformation.
 *
 * Moves are in the form: [{x, y, color}, ...]
 *
 * @param {string} board - rectangular grid of characters 1-5 separated by newlines
 * @param {Array<{x: number, y: number, color: number}>} moves - array of move objects
 * @returns {string} the board after all moves have been applied
 */
function floodFillPlayer(board, moves) {
}

module.exports = floodFillPlayer;