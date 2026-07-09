const fix = f => (g => g(g))(x => f(y => x(x)(y)));

const factorial = f => n => n <= 1n ? 1n : n * f(n - 1n);

const fibonacci = f => n => n <= 1 ? BigInt(n) : f(n - 1) + f(n - 2);

const foldr = self => folding_fn => init => gen => {
  const result = gen.next();
  if (result.done) {
    return init;
  }
  return folding_fn(result.value)(() => self(folding_fn)(init)(gen));
};