function rot(s) {
  return s.split('\n').map(l => l.split('').reverse().join('')).reverse().join('\n');
}

function selfieAndRot(s) {
  const lines = s.split('\n');
  const n = lines.length;
  const dots = '.'.repeat(n);
  const top = lines.map(l => l + dots);
  const bot = rot(s).split('\n').map(l => dots + l);
  return top.concat(bot).join('\n');
}

function oper(fct, s) {
  return fct(s);
}