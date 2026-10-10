function getAge(birthDate, nowDate = new Date()) {
  let age = nowDate.getFullYear() - birthDate.getFullYear();
  
  const birthMonth = birthDate.getMonth();
  const birthDay = birthDate.getDate();
  const nowMonth = nowDate.getMonth();
  const nowDay = nowDate.getDate();
  
  if (nowMonth < birthMonth || (nowMonth === birthMonth && nowDay < birthDay)) {
    age--;
  }
  
  return age;
}