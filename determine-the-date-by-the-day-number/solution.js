function dateFromDay(day, isLeap) {
  const months = ['January', 'February', 'March', 'April', 'May', 'June', 
                  'July', 'August', 'September', 'October', 'November', 'December'];
  const daysInMonth = [31, isLeap ? 29 : 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];
  
  let remainingDays = day;
  for (let i = 0; i < 12; i++) {
    if (remainingDays <= daysInMonth[i]) {
      return `${months[i]}, ${remainingDays}`;
    }
    remainingDays -= daysInMonth[i];
  }
}