function countDeafRats(town) {
  town = town.replace(/ /g, '');
  const p = town.indexOf('P');

  // rats to the left should head right (~O); to the right should head left (O~)
  const left = (town.slice(0, p).match(/O~|~O/g) || []);
  const right = (town.slice(p + 1).match(/O~|~O/g) || []);

  let deaf = 0;
  for (const r of left) if (r === 'O~') deaf++;
  for (const r of right) if (r === '~O') deaf++;

  return deaf;
}