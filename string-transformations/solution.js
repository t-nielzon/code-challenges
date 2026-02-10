function stringTransformation(s, transformations) {
  let result = s;
  for (const op of transformations) {
    switch (op) {
      case "U":
        result = result.toUpperCase();
        break;
      case "L":
        result = result.toLowerCase();
        break;
      case "R":
        result = result.split("").reverse().join("");
        break;
      case "D":
        result = result.replace(/./g, "$&$&");
        break;
    }
  }
  return result;
}

module.exports = { stringTransformation };