function correct(exp) {
  const [lhs, rhsStr] = exp.split('=').map(s => s.trim());
  const target = parseFloat(rhsStr);
  const ops = ['+', '-', '*', '/'];
  const slots = (lhs.match(/\(\)/g) || []).length;

  const tryCombo = (idx, current) => {
    if (idx === slots) {
      try {
        const val = eval(current);
        if (Math.abs(val - target) < 1e-9) return current;
      } catch (e) {}
      return null;
    }
    for (const op of ops) {
      const next = current.replace('()', `(${op})`);
      const result = tryCombo(idx + 1, next);
      if (result !== null) return result;
    }
    return null;
  };

  const solvedLhs = tryCombo(0, lhs);
  return `${solvedLhs} = ${rhsStr}`;
}