function validate(str) {
  const regex = /^(?!.*(name:.*){2})(?!.*(age:.*){2})(?!.*(skill:.*){2})(?=.*name:[A-Za-z]+)(?=.*age:[0-9]+)(?=.*skill:[A-Za-z]+)(name:[A-Za-z]+|age:[0-9]+|skill:[A-Za-z]+)(,(name:[A-Za-z]+|age:[0-9]+|skill:[A-Za-z]+)){2}$/;
  return regex.test(str);
}