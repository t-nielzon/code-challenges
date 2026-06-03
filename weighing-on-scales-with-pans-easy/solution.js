/*
 * Weighing on scales with pans (easy)
 *
 * Balance condition: left pan == right pan + N.
 * Each piece is assigned to the left pan (+1), the right pan (-1), or left
 * unused (0). We need the signed sum to equal N. With at most 10 pieces, a
 * full 3^n enumeration is cheap.
 */

function weighing(weights, N) {
  const n = weights.length;

  // assign[i] in {1: left, -1: right, 0: unused}
  function search(i, balance, left, right) {
    if (i === n) {
      // left total must exceed right total by exactly N
      return balance === N ? [left, right] : null;
    }

    const w = weights[i];

    // put on left pan
    let res = search(i + 1, balance + w, [...left, w], right);
    if (res) return res;

    // put on right pan
    res = search(i + 1, balance - w, left, [...right, w]);
    if (res) return res;

    // leave unused
    return search(i + 1, balance, left, right);
  }

  return search(0, 0, [], []) || [];
}

module.exports = weighing;