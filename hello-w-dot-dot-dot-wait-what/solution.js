function greeting() {
  const a = +!![];        // 1
  const b = a + a;        // 2
  const c = b + b;        // 4
  const d = c + c;        // 8
  const e = d + d;        // 16
  const f = e + e;        // 32
  const g = f + f;        // 64

  return String.fromCharCode(
    g + d,                 // 72  H
    g + f + c + a,         // 101 e
    g + f + d + c,         // 108 l
    g + f + d + c,         // 108 l
    g + f + d + c + b + a, // 111 o
    f,                     // 32  (space)
    g + e + c + b + a,     // 87  W
    g + f + d + c + b + a, // 111 o
    g + f + e + b,         // 114 r
    g + f + d + c,         // 108 l
    g + f + c,             // 100 d
    f + a                  // 33  !
  );
}