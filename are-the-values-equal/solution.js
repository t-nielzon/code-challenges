function areEqual(a, b) {
  const values = (x) => {
    if (typeof x === 'string') {
      return x.split('').map((c) => JSON.stringify(c)).sort();
    }
    if (Array.isArray(x)) {
      return x.map((v) => canonical(v)).sort();
    }
    if (x !== null && typeof x === 'object') {
      return Object.keys(x)
        .sort()
        .map((k) => k + ':' + canonical(x[k]));
    }
    return [canonical(x)];
  };

  const canonical = (x) => {
    if (x !== null && typeof x === 'object') {
      if (Array.isArray(x)) {
        return '[' + x.map((v) => canonical(v)).sort().join(',') + ']';
      }
      return (
        '{' +
        Object.keys(x)
          .sort()
          .map((k) => JSON.stringify(k) + ':' + canonical(x[k]))
          .join(',') +
        '}'
      );
    }
    return JSON.stringify(x);
  };

  const sa = values(a);
  const sb = values(b);
  if (sa.length !== sb.length) return false;
  return sa.every((v, i) => v === sb[i]);
}