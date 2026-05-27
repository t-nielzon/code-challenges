function solution(str) {
  if (str.length % 2 !== 0) str += '_';
  return str.match(/.{2}/g) || [];
}