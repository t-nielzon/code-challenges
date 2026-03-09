const md5 = require('js-md5');

function crack(hash) {
  for (let i = 0; i <= 99999; i++) {
    const pin = String(i).padStart(5, '0');
    if (md5(pin) === hash) return pin;
  }
}