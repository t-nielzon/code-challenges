function chess960(i) {
  const pieces = ['B','B','K','N','N','Q','R','R'];
  const valid = [];

  function permute(arr, l) {
    if (l === arr.length) {
      const b1 = arr.indexOf('B');
      const b2 = arr.lastIndexOf('B');
      if (b1 % 2 === b2 % 2) return;
      const k = arr.indexOf('K');
      const r1 = arr.indexOf('R');
      const r2 = arr.lastIndexOf('R');
      if (!(r1 < k && k < r2)) return;
      valid.push(arr.join(''));
      return;
    }
    const used = new Set();
    for (let j = l; j < arr.length; j++) {
      if (used.has(arr[j])) continue;
      used.add(arr[j]);
      [arr[l], arr[j]] = [arr[j], arr[l]];
      permute(arr, l + 1);
      [arr[l], arr[j]] = [arr[j], arr[l]];
    }
  }

  permute(pieces, 0);
  valid.sort();

  const rank = valid[i - 1];
  const black = rank.toLowerCase().split('').join(' ');
  const white = rank.split('').join(' ');
  const e = '. . . . . . . .';
  const bp = 'p p p p p p p p';
  const wp = 'P P P P P P P P';

  return `${black}\n${bp}\n${e}\n${e}\n${e}\n${e}\n${wp}\n${white}\n`;
}