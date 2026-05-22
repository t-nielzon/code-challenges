function countDeafRats(town) {
  const s = town.replace(/ /g, '');
  const piperIndex = s.indexOf('P');
  let deaf = 0;

  for (let i = 0; i < s.length; i += 2) {
    if (i === piperIndex) {
      i -= 1;
      continue;
    }
    const rat = s.substr(i, 2);
    if (i < piperIndex && rat === 'O~') deaf++;
    if (i > piperIndex && rat === '~O') deaf++;
  }

  return deaf;
}