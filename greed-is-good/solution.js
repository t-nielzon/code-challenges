function score(dice) {
  const counts = [0, 0, 0, 0, 0, 0, 0];
  for (const d of dice) counts[d]++;

  const tripleScore = { 1: 1000, 2: 200, 3: 300, 4: 400, 5: 500, 6: 600 };
  const singleScore = { 1: 100, 5: 50 };

  let total = 0;
  for (let face = 1; face <= 6; face++) {
    const c = counts[face];
    if (c >= 3) total += tripleScore[face];
    const remaining = c >= 3 ? c - 3 : c;
    if (singleScore[face]) total += remaining * singleScore[face];
  }
  return total;
}