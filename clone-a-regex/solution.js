RegExp.clone = function (regex) {
  const cloned = new RegExp(regex.source, regex.flags);
  cloned.lastIndex = regex.lastIndex;
  return cloned;
};

RegExp.prototype.clone = function () {
  return RegExp.clone(this);
};