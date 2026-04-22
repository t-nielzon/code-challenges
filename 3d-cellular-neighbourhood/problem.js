/*
 * 3D Cellular Neighbourhood
 *
 * Given a neighbourhood type ("moore" or "von_neumann"), a MxNxK 3D matrix
 * (list of lists of lists), a 3-tuple of coordinates and the distance,
 * return the list of neighbours of the given cell.
 *
 * Order of the indices: The first index should be applied for the outer/first
 * matrix layer. The last index for the most inner/last layer.
 * coordinates = (i, j, k) should be applied like mat[i][j][k]
 *
 * Return an empty array if any of these conditions are true:
 *  - Matrix is empty
 *  - Coordinates are outside the matrix
 *  - Distance is equal to 0
 *
 * Moore neighborhood: cells that shape a 'cube' around the given cell.
 * Von Neumann neighborhood: cells that shape an 'octahedron' (diamond) around the given cell.
 *
 * A cell is not its own neighbor.
 */

function getNeighbourhood(type, mat, coordinates, distance) {
  // your code here
}