function productFib(prod) {
  let fn = 0;
  let fn1 = 1;
  
  while (fn * fn1 < prod) {
    let temp = fn + fn1;
    fn = fn1;
    fn1 = temp;
  }
  
  return [fn, fn1, fn * fn1 === prod];
}