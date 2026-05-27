/*
 * As a historian, you wonder how many people ever lived in a specific area.
 *
 * Given an amount of records as an array of arrays representing the year of
 * arrival/birth and departure/death from a given area, write a function that
 * returns the maximum for that area and the year in which said maximum was
 * achieved.
 *
 * If more years have that maximum, just return the first one of them.
 *
 * If you have both people entering and exiting the territory in the same year,
 * consider that all new comers arrived/were born before anyone left/died for
 * the sake of computing when the maximum was achieved.
 *
 * You will always be given at least one record of beginning/ending of a
 * presence on the territory in valid format (ie: the first value will be
 * lesser than or equal to the second).
 *
 * Examples:
 *   yearMaxPeople([[1978, 1978], [1969, 1998]]) === [2, 1978]
 *   yearMaxPeople([[1980, 2010], [1979, 1985], [1986, 1995], [1987, 2008]]) === [3, 1987]
 *   yearMaxPeople([[1980, 2010], [1979, 1986], [1986, 1995], [1987, 2008]]) === [3, 1986]
 */

function yearMaxPeople(people) {

}