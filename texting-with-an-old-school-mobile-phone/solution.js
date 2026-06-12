function sendMessage(message) {
  const PRESS = {
    '.': '1', ',': '11', '?': '111', '!': '1111',
    a: '2', b: '22', c: '222',
    d: '3', e: '33', f: '333',
    g: '4', h: '44', i: '444',
    j: '5', k: '55', l: '555',
    m: '6', n: '66', o: '666',
    p: '7', q: '77', r: '777', s: '7777',
    t: '8', u: '88', v: '888',
    w: '9', x: '99', y: '999', z: '9999',
    "'": '*', '-': '**', '+': '***', '=': '****',
    ' ': '0',
  };

  let result = '';
  let upper = false;
  let lastKey = '';

  for (const ch of message) {
    let token;
    if (/[0-9*#]/.test(ch)) {
      // digits and *# are typed by holding: key followed by a dash.
      // the dash also means the next character never needs a wait.
      token = ch + '-';
    } else {
      token = '';
      const lower = ch.toLowerCase();
      if (/[a-z]/.test(lower)) {
        const isUpper = ch !== lower;
        if (isUpper !== upper) {
          // toggle case only when an alphabetic character requires it
          token += '#';
          upper = isUpper;
        }
      }
      token += PRESS[lower];
    }
    // wait (space) only when this token starts by pressing the same key
    // the previous token ended on; a leading '#' (case toggle) or a
    // previous trailing '-' (hold) breaks the sequence, so no wait.
    if (lastKey && token[0] === lastKey) result += ' ';
    result += token;
    lastKey = token[token.length - 1];
  }

  return result;
}