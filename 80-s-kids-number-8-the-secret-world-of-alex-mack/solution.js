function findSecret(employees) {
  return employees.find(e => e.lastName === 'Mack').badgeNumber;
}