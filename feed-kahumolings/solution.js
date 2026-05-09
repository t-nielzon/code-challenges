function feed(food, flavour, mouths) {
  let first;
  if (flavour === 1) {
    first = food / mouths;
  } else {
    first = food * (flavour - 1) / (Math.pow(flavour, mouths) - 1);
  }
  const result = [];
  let current = first;
  for (let i = 0; i < mouths; i++) {
    result.push(Math.round(current * 1000) / 1000);
    current *= flavour;
  }
  return result;
}