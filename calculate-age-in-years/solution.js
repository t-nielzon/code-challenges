function getAge(birthDate, nowDate = new Date()) {
  let age = nowDate.getFullYear() - birthDate.getFullYear();
  
  const currentMonth = nowDate.getMonth();
  const birthMonth = birthDate.getMonth();
  const currentDay = nowDate.getDate();
  const birthDay = birthDate.getDate();
  
  // subtract a year if birthday hasn't occurred yet this year
  if (currentMonth < birthMonth || (currentMonth === birthMonth && currentDay < birthDay)) {
    age--;
  }
  
  return age;
}