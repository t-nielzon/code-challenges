function prettyfy(source) {
  const tokens = [];
  const re = /<[^>]+>/g;
  let last = 0, m;
  while ((m = re.exec(source)) !== null) {
    if (m.index > last) tokens.push({ type: 'text', value: source.slice(last, m.index) });
    tokens.push({ type: 'tag', value: m[0] });
    last = re.lastIndex;
  }
  if (last < source.length) tokens.push({ type: 'text', value: source.slice(last) });

  // collapse whitespace in text, drop pure-whitespace runs
  const clean = [];
  for (const t of tokens) {
    if (t.type === 'text') {
      const v = t.value.replace(/\s+/g, ' ').trim();
      if (v) clean.push({ type: 'text', value: v });
    } else {
      clean.push(t);
    }
  }

  let level = 0;
  const lines = [];
  let prevType = null;

  for (const t of clean) {
    if (t.type === 'text') {
      lines.push('  '.repeat(level) + t.value);
      prevType = 'text';
      continue;
    }

    const tag = t.value;
    const isSelf = tag.endsWith('/>');
    const isClose = !isSelf && tag.startsWith('</');
    const isBr = /^<br\s*\/>$/.test(tag);

    if (isBr && prevType === 'text') {
      // br after text attaches to the text line, then a newline follows
      lines[lines.length - 1] += tag;
    } else if (isSelf) {
      lines.push('  '.repeat(level) + tag);
    } else if (isClose) {
      level--;
      lines.push('  '.repeat(level) + tag);
    } else {
      lines.push('  '.repeat(level) + tag);
      level++;
    }
    prevType = 'tag';
  }

  return lines.join('\n') + '\n';
}