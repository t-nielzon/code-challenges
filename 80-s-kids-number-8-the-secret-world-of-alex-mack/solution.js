/*
 * Alex Mack's father in the show is George Mack.
 * Find his employee record and return his badge number.
 */
function findSecret(employees) {
  return employees.find(e => e.lastName === 'Mack' && e.firstName === 'George').badgeNumber;
}