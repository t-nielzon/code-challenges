function reverseStr(s) {
  return s.split('').reverse().join('');
}

function smallestS(P) {
  const len = P.length;
  if (len === 0) return null;

  let zeros = 0;
  let currentLen = len;
  while (currentLen > 1) {
    let allNines = true;
    for (let i = 0; i < currentLen - 1; i++) {
      if (P[i] !== '9') { allNines = false; break; }
    }
    if (allNines) break;
    zeros++;
    currentLen--;
  }

  let tail;
  if (currentLen === 1) {
    if (P[0] === '9') return null;
    tail = String(parseInt(P[0], 10) + 1);
  } else {
    if (P[currentLen - 1] === '9') return null;
    tail = String(parseInt(P[currentLen - 1], 10) + 1) + '9'.repeat(currentLen - 1);
  }

  return '0'.repeat(zeros) + tail;
}

function doublyNotLess(n) {
  const L = n.length;
  if (reverseStr(n) >= n) return n;

  let best = null;

  for (let k = L - 1; k >= 0; k--) {
    const prefix = n.slice(0, k);
    const nkChar = n[k];
    if (nkChar === '9') continue;
    const nk = parseInt(nkChar, 10);

    const P = n.slice(0, L - 1 - k);
    const Q = n.slice(L - 1 - k);
    const revPrefix = reverseStr(prefix);

    const s_ii = smallestS(P);
    if (s_ii !== null) {
      const m_ii = prefix + String(nk + 1) + s_ii;
      if (best === null || m_ii < best) best = m_ii;
    }

    const s_i = reverseStr(P);
    for (let c = nk + 1; c <= 9; c++) {
      const U = String(c) + revPrefix;
      if (U >= Q) {
        const m_i = prefix + String(c) + s_i;
        if (best === null || m_i < best) best = m_i;
        break;
      }
    }
  }

  if (best === null) {
    best = '1' + '0'.repeat(L - 1) + '1';
  }

  return best;
}