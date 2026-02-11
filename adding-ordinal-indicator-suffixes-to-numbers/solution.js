function numberToOrdinal(n) {
  if (n === 0) return "0";

  const lastTwo = n % 100;

  if (lastTwo >= 11 && lastTwo <= 13) return n + "th";

  switch (n % 10) {
    case 1: return n + "st";
    case 2: return n + "nd";
    case 3: return n + "rd";
    default: return n + "th";
  }
}

module.exports = { numberToOrdinal };