function amidakuji(ladder) {
  const width = ladder[0].length + 1;
  let positions = Array.from({ length: width }, (_, i) => i);

  for (const row of ladder) {
    for (let i = 0; i < row.length; i++) {
      if (row[i] === '1') {
        [positions[i], positions[i + 1]] = [positions[i + 1], positions[i]];
      }
    }
  }

  return positions;
}