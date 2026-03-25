function numberToEnglish(n) {
  if (!Number.isInteger(n) || n < 0 || n > 99999) return '';

  const ones = ['zero','one','two','three','four','five','six','seven','eight','nine',
    'ten','eleven','twelve','thirteen','fourteen','fifteen','sixteen','seventeen','eighteen','nineteen'];
  const tens = ['','','twenty','thirty','forty','fifty','sixty','seventy','eighty','ninety'];

  if (n === 0) return 'zero';

  function underHundred(num) {
    if (num === 0) return '';
    if (num < 20) return ones[num];
    const t = tens[Math.floor(num / 10)];
    const o = num % 10 ? ' ' + ones[num % 10] : '';
    return t + o;
  }

  function underThousand(num) {
    if (num === 0) return '';
    if (num < 100) return underHundred(num);
    const h = ones[Math.floor(num / 100)] + ' hundred';
    const rem = num % 100;
    return rem ? h + ' ' + underHundred(rem) : h;
  }

  const thousands = Math.floor(n / 1000);
  const remainder = n % 1000;

  let result = '';
  if (thousands) result += underThousand(thousands) + ' thousand';
  if (remainder) result += (thousands ? ' ' : '') + underThousand(remainder);

  return result;
}