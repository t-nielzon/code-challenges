function ann(n) {
  if (n === 0) return [];
  
  const a = [1];
  const j = [0];
  
  for (let i = 1; i < n; i++) {
    const jVal = i - a[j[i-1]];
    const aVal = i - j[a[i-1]];
    j.push(jVal);
    a.push(aVal);
  }
  
  return a;
}

function john(n) {
  if (n === 0) return [];
  
  const a = [1];
  const j = [0];
  
  for (let i = 1; i < n; i++) {
    const jVal = i - a[j[i-1]];
    const aVal = i - j[a[i-1]];
    j.push(jVal);
    a.push(aVal);
  }
  
  return j;
}

function sum_ann(n) {
  if (n === 0) return 0;
  
  const a = [1];
  const j = [0];
  let sum = 1;
  
  for (let i = 1; i < n; i++) {
    const jVal = i - a[j[i-1]];
    const aVal = i - j[a[i-1]];
    j.push(jVal);
    a.push(aVal);
    sum += aVal;
  }
  
  return sum;
}

function sum_john(n) {
  if (n === 0) return 0;
  
  const a = [1];
  const j = [0];
  let sum = 0;
  
  for (let i = 1; i < n; i++) {
    const jVal = i - a[j[i-1]];
    const aVal = i - j[a[i-1]];
    j.push(jVal);
    a.push(aVal);
    sum += jVal;
  }
  
  return sum;
}