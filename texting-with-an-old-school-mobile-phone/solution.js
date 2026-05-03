function sendMessage(message) {
  const letterMap = {
    a: '2', b: '22', c: '222',
    d: '3', e: '33', f: '333',
    g: '4', h: '44', i: '444',
    j: '5', k: '55', l: '555',
    m: '6', n: '66', o: '666',
    p: '7', q: '77', r: '777', s: '7777',
    t: '8', u: '88', v: '888',
    w: '9', x: '99', y: '999', z: '9999'
  };
  const specialMap = {
    '.': '1', ',': '11', '?': '111', '!': '1111',
    "'": '*', '-': '**', '+': '***', '=': '****',
    ' ': '0'
  };
  const holdChars = '0123456789*#';

  let result = '';
  let upper = false;
  let lastKey = null;
  let lastWasHold = false;

  for (const ch of message) {
    let seq;
    let isHold = false;
    let caseSwitched = false;

    const lower = ch.toLowerCase();
    if (letterMap[lower]) {
      const isUpper = ch !== lower;
      if (isUpper !== upper) {
        result += '#';
        upper = isUpper;
        caseSwitched = true;
      }
      seq = letterMap[lower];
    } else if (specialMap[ch] !== undefined) {
      seq = specialMap[ch];
    } else if (holdChars.includes(ch)) {
      seq = ch + '-';
      isHold = true;
    } else {
      continue;
    }

    const key = seq[0];
    if (lastKey === key && !lastWasHold && !caseSwitched) {
      result += ' ';
    }
    result += seq;
    lastKey = key;
    lastWasHold = isHold;
  }

  return result;
}