function getDate(day, isLeap) {
  const daysInMonths = isLeap 
    ? [31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    : [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
  
  const monthNames = ["January", "February", "March", "April", "May", "June",
                      "July", "August", "September", "October", "November", "December"];
  
  let remaining = day;
  for (let i = 0; i < 12; i++) {
    if (remaining <= daysInMonths[i]) {
      return `${monthNames[i]}, ${remaining}`;
    }
    remaining -= daysInMonths[i];
  }
}