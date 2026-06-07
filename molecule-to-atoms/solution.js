function parseMolecule(formula) {
  const tokens = formula.match(/[A-Z][a-z]*|\d+|[()\[\]{}]/g) || [];
  let pos = 0;

  const opening = { '(': ')', '[': ']', '{': '}' };
  const closing = { ')': '(', ']': '[', '}': '{' };

  function readNumber() {
    if (pos < tokens.length && /^\d+$/.test(tokens[pos])) {
      return parseInt(tokens[pos++], 10);
    }
    return 1;
  }

  function parseGroup() {
    const counts = {};

    function add(element, n) {
      counts[element] = (counts[element] || 0) + n;
    }

    while (pos < tokens.length) {
      const token = tokens[pos];

      if (closing[token]) {
        break;
      } else if (opening[token]) {
        pos++; // consume opening bracket
        const inner = parseGroup();
        const close = tokens[pos++]; // consume closing bracket
        if (!close || opening[token] !== close) {
          throw new Error('Not a valid molecule');
        }
        const multiplier = readNumber();
        for (const element in inner) {
          add(element, inner[element] * multiplier);
        }
      } else if (/^[A-Z][a-z]*$/.test(token)) {
        pos++; // consume element
        const n = readNumber();
        add(token, n);
      } else {
        throw new Error('Not a valid molecule');
      }
    }

    return counts;
  }

  const result = parseGroup();
  if (pos !== tokens.length) {
    throw new Error('Not a valid molecule');
  }
  return result;
}