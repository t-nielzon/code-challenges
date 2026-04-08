function Arith(num) {
  this.value = toNumber(num);
}

Arith.prototype.add = function(num) {
  return toWords(this.value + toNumber(num));
};

const ones = ['zero','one','two','three','four','five','six','seven','eight','nine',
  'ten','eleven','twelve','thirteen','fourteen','fifteen','sixteen','seventeen','eighteen','nineteen'];
const tens = ['','','twenty','thirty','forty','fifty','sixty','seventy','eighty','ninety'];

function toNumber(words) {
  if (words === 'zero') return 0;
  const parts = words.split(' ').filter(w => w !== 'and');
  let result = 0;
  let i = 0;
  while (i < parts.length) {
    if (parts[i + 1] === 'hundred') {
      result += wordToNum(parts[i]) * 100;
      i += 2;
    } else {
      result += wordToNum(parts[i]);
      i++;
    }
  }
  return result;
}

function wordToNum(w) {
  let idx = ones.indexOf(w);
  if (idx !== -1) return idx;
  idx = tens.indexOf(w);
  if (idx !== -1) return idx * 10;
  return 0;
}

function toWords(n) {
  if (n === 0) return 'zero';
  if (n === 1000) return 'one thousand';
  const parts = [];
  if (n >= 100) {
    parts.push(ones[Math.floor(n / 100)] + ' hundred');
    n %= 100;
    if (n > 0) parts.push('and');
  }
  if (n >= 20) {
    let s = tens[Math.floor(n / 10)];
    if (n % 10 !== 0) s += ' ' + ones[n % 10];
    parts.push(s);
  } else if (n > 0) {
    parts.push(ones[n]);
  }
  return parts.join(' ');
}