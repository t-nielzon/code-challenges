function iterPi(epsilon) {
  let pi4 = 0;
  let i = 0;
  let sign = 1;

  while (true) {
    pi4 += sign / (2 * i + 1);
    sign = -sign;
    i++;
    if (Math.abs(4 * pi4 - Math.PI) < epsilon) {
      return [i, Math.round(4 * pi4 * 1e10) / 1e10];
    }
  }
}