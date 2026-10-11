function countFridayMonths(year1, year2) {
  let count = 0;
  
  for (let year = year1; year <= year2; year++) {
    for (let month = 1; month <= 12; month++) {
      // get the last day of the month
      const lastDay = new Date(year, month, 0);
      
      // check if it's a friday (day 5)
      if (lastDay.getDay() === 5) {
        count++;
      }
    }
  }
  
  return count;
}