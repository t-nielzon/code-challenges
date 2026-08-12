/*
In lambda calculus, the only primitives are lambdas. This kata requires implementing
boolean operators (not, and, or) using Church encoding, but with an unusual twist:
the operators are applied in postfix and infix notation using single-letter identifiers.

Implement the following operators:
- not: Applied postfix as b(n)(o)(t)
- and: Applied infix as b1(a)(n)(d)(b2)
- or: Applied infix as b1(o)(r)(b2)

The Church boolean definitions are provided:
- True = t => f => t (returns first argument)
- False = t => f => f (returns second argument)

Your task is to define the operators n, o, t, a, d, r such that:
- True(n)(o)(t) == False  (not True = False)
- False(n)(o)(t) == True  (not False = True)
- True(a)(n)(d)(True) == True  (True and True = True)
- True(a)(n)(d)(False) == False  (True and False = False)
- False(a)(n)(d)(x) == False  (False and x = False)
- False(o)(r)(False) == False  (False or False = False)
- False(o)(r)(True) == True  (False or True = True)
- True(o)(r)(x) == True  (True or x = True)
*/

function solution() {
  const False = t => f => f;
  const True = t => f => t;

  // Implement your operators here
  const n = /* TODO */;
  const o = /* TODO */;
  const t = /* TODO */;
  const a = /* TODO */;
  const d = /* TODO */;
  const r = /* TODO */;

  return { n, o, t, a, d, r };
}