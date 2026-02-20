function SubstitutionCipher(abc1, abc2) {
  this.abc1 = abc1;
  this.abc2 = abc2;
}

SubstitutionCipher.prototype.encode = function (str) {
  return [...str].map(c => {
    const i = this.abc1.indexOf(c);
    return i === -1 ? c : this.abc2[i];
  }).join('');
};

SubstitutionCipher.prototype.decode = function (str) {
  return [...str].map(c => {
    const i = this.abc2.indexOf(c);
    return i === -1 ? c : this.abc1[i];
  }).join('');
};