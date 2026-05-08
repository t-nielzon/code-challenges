function score(dice) {
  const counts = [0, 0, 0, 0, 0, 0, 0];
  for (const d of dice) counts[d]++;

  const tripletPoints = { 1: 1000, 2: 200, 3: 300, 4: 400, 5: 500, 6: 600 };
  let total = 0;

  for (let face = 1; face <= 6; face++) {
    const c = counts[face];
    const triples = Math.floor(c / 3);
    const remainder = c % 3;
    total += triples * tripletPoints[face];
    if (face === 1) total += remainder * 100;
    else if (face === 5) total += remainder * 50;
  }

  return total;
}