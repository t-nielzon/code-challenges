function numberToOrdinal(n) {
  if (n === 0) return "0";
  
  const lastTwo = n % 100;
  const lastOne = n % 10;
  
  // Handle the teen exceptions (11, 12, 13)
  if (lastTwo >= 11 && lastTwo <= 13) {
    return n + "th";
  }
  
  // Handle regular cases based on last digit
  switch (lastOne) {
    case 1:
      return n + "st";
    case 2:
      return n + "nd";
    case 3:
      return n + "rd";
    default:
      return n + "th";
  }
}