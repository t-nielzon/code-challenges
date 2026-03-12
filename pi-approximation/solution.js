function iterPi(epsilon) {
  let pi4 = 0;
  let i = 0;
  let sign = 1;

  while (true) {
    pi4 += sign / (2 * i + 1);
    sign = -sign;
    i++;
    const pi = pi4 * 4;
    if (Math.abs(pi - Math.PI) < epsilon) {
      return [i, Math.round(pi * 1e10) / 1e10];
    }
  }
}

module.exports = { iterPi };