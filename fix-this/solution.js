const fix = f => (x => f(y => x(x)(y)))(x => f(y => x(x)(y)));

const factorial = f => n => n <= 1n ? 1n : n * f(n - 1n);

const fibonacci = f => n => n <= 1 ? BigInt(n) : f(n - 1) + f(n - 2);

const foldr = rec => f => acc => gen => {
  const next = gen.next();
  if (next.done) {
    return acc;
  }
  return f(next.value)(() => rec(f)(acc)(gen));
};