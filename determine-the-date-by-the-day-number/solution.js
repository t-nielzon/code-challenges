function dateFromDayNumber(day, isLeap) {
  const monthNames = ["January", "February", "March", "April", "May", "June",
                      "July", "August", "September", "October", "November", "December"];
  const daysInMonth = [31, isLeap ? 29 : 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
  
  let dayCount = 0;
  for (let i = 0; i < 12; i++) {
    if (dayCount + daysInMonth[i] >= day) {
      return `${monthNames[i]}, ${day - dayCount}`;
    }
    dayCount += daysInMonth[i];
  }
}