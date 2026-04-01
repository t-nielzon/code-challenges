function findSecret(employees) {
  var found = employees.find(function(emp) {
    return emp.lastName === 'Mack';
  });
  return found.badgeNumber;
}