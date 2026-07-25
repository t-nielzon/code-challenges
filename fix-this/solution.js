function fix(f) {
  return f(x => fix(f)(x));
}

function factorial(f) {
  return n => n <= 1n ? 1n : n * f(n - 1n);
}

function fibonacci(f) {
  return n => n <= 1n ? BigInt(n) : f(n - 1n) + f(n - 2n);
}

function foldr(f) {
  return z => gen => {
    const {value, done} = gen.next();
    if (done) return z;
    return f(value)(() => foldr(f)(z)(gen));
  };
}