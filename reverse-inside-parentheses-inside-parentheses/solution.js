function reverseInParentheses(str) {
  const chars = str.split('');
  const depths = new Array(chars.length).fill(0);
  let depth = 0;
  for (let i = 0; i < chars.length; i++) {
    if (chars[i] === '(') {
      depth++;
      depths[i] = depth;
    } else if (chars[i] === ')') {
      depths[i] = depth;
      depth--;
    } else {
      depths[i] = depth;
    }
  }

  function recurse(start, end, level) {
    const segment = chars.slice(start, end);
    const groups = [];
    let i = 0;
    while (i < segment.length) {
      const absIdx = start + i;
      if (depths[absIdx] > level && (segment[i] === '(' || depths[absIdx] >= level + 1)) {
        let j = i;
        let d = 0;
        while (j < segment.length) {
          if (segment[j] === '(') d++;
          else if (segment[j] === ')') {
            d--;
            if (d === 0) { j++; break; }
          }
          j++;
        }
        const inner = recurse(start + i + 1, start + j - 1, level + 1);
        groups.push({ type: 'paren', content: inner });
        i = j;
      } else {
        let j = i;
        while (j < segment.length) {
          const a = start + j;
          if (segment[j] === '(' && depths[a] > level) break;
          j++;
        }
        groups.push({ type: 'text', content: segment.slice(i, j).join('') });
        i = j;
      }
    }

    if (level % 2 === 1) {
      groups.reverse();
      return groups.map(g => {
        if (g.type === 'text') return g.content.split('').reverse().join('');
        return ')' + g.content + '(';
      }).join('');
    } else {
      return groups.map(g => {
        if (g.type === 'text') return g.content;
        return '(' + g.content + ')';
      }).join('');
    }
  }

  return recurse(0, chars.length, 0);
}

module.exports = reverseInParentheses;