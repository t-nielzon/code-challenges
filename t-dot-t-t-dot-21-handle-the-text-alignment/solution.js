function formatScoreList(list) {
  const avg = list.reduce((s, x) => s + x.score, 0) / list.length;
  const avgStr = avg.toFixed(1);

  const leftItems = list.map(x => x.name).concat(['The average score']);
  const leftWidth = Math.max(...leftItems.map(s => s.length));

  const rightItems = list.map(x => x.score.toString() + '  pts').concat([avgStr + ' pts']);
  const rightWidth = Math.max(...rightItems.map(s => s.length));

  const sep = '+' + '-'.repeat(leftWidth + 1) + '+' + '-'.repeat(rightWidth + 1) + '+';

  // center "Name" in leftWidth+1 wide cell, with padding bias to right
  const nameLabel = 'Name';
  const nameSpace = leftWidth + 1 - nameLabel.length;
  const nameLeft = Math.floor(nameSpace / 2);
  const nameRight = nameSpace - nameLeft;
  const nameCentered = ' '.repeat(nameLeft) + nameLabel + ' '.repeat(nameRight);

  // center "Score" in rightWidth+1 wide cell
  const scoreLabel = 'Score';
  const scoreSpace = rightWidth + 1 - scoreLabel.length;
  const scoreLeft = Math.floor(scoreSpace / 2);
  const scoreRight = scoreSpace - scoreLeft;
  const scoreCentered = ' '.repeat(scoreLeft) + scoreLabel + ' '.repeat(scoreRight);

  const header = '|' + nameCentered + '|' + scoreCentered + '|';

  const rows = [];
  rows.push(sep);
  rows.push(header);
  rows.push(sep);

  for (const item of list) {
    const left = item.name.padEnd(leftWidth + 1);
    const scoreStr = item.score.toString();
    // right part: score flush left, "pts" flush right, padded with spaces
    const rightContent = scoreStr + ' '.repeat(rightWidth - scoreStr.length - 3) + 'pts';
    rows.push('|' + left + '|' + rightContent + '|');
    rows.push(sep);
  }

  // average row
  const avgLeft = 'The average score'.padEnd(leftWidth + 1);
  const avgRight = avgStr + ' '.repeat(rightWidth - avgStr.length - 3) + 'pts';
  rows.push('|' + avgLeft + '|' + avgRight + '|');
  rows.push(sep);

  return rows.join('\n');
}