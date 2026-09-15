function* fibonacci() {
  let a = 1n;
  let b = 1n;
  
  yield a;
  yield b;
  
  while (true) {
    const next = a + b;
    yield next;
    a = b;
    b = next;
  }
}