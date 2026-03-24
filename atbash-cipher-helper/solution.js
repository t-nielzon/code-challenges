class AtbashCipher {
  constructor(alphabet) {
    this.map = {};
    for (let i = 0; i < alphabet.length; i++) {
      this.map[alphabet[i]] = alphabet[alphabet.length - 1 - i];
    }
  }

  encode(string) {
    return [...string].map(c => this.map[c] ?? c).join('');
  }

  decode(string) {
    return this.encode(string);
  }
}