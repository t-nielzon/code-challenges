function getAge(birthDate, now = new Date()) {
  let age = now.getFullYear() - birthDate.getFullYear();
  
  const birthdayThisYear = new Date(
    now.getFullYear(),
    birthDate.getMonth(),
    birthDate.getDate()
  );
  
  if (now < birthdayThisYear) {
    age--;
  }
  
  return age;
}