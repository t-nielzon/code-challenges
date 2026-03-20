var Morse = (function() {
  var morseTable = {
    'A':'.-','B':'-...','C':'-.-.','D':'-..','E':'.','F':'..-.','G':'--.','H':'....',
    'I':'..','J':'.---','K':'-.-','L':'.-..','M':'--','N':'-.','O':'---','P':'.--.',
    'Q':'--.-','R':'.-.','S':'...','T':'-','U':'..-','V':'...-','W':'.--','X':'-..-',
    'Y':'-.--','Z':'--..','0':'-----','1':'.----','2':'..---','3':'...--','4':'....-',
    '5':'.....','6':'-....','7':'--...','8':'---..','9':'----.','.'  :'.-.-.-',
    ',':'--..--','?':'..--..','\'':'.----.','!':'-.-.--','/':'-..-.','(':'-.--.',
    ')':'-.--.-','&':'.-...',':':'---...',';':'-.-.-.','=':'-...-','+':'.-.-.',
    '-':'-....-','_':'..--.−','"':'.-..-.','$':'...-..-','@':'.--.-.'
  };

  // fix underscore and rebuild
  morseTable['_'] = '..--.-';

  var reverseMorse = {};
  for (var ch in morseTable) {
    reverseMorse[morseTable[ch]] = ch;
  }

  function charToBits(c) {
    var morse = morseTable[c.toUpperCase()];
    if (!morse) return '';
    var parts = [];
    for (var i = 0; i < morse.length; i++) {
      parts.push(morse[i] === '.' ? '1' : '111');
    }
    return parts.join('0');
  }

  function messageToBitString(message) {
    var words = message.split(' ');
    var wordBits = words.map(function(word) {
      var chars = [];
      for (var i = 0; i < word.length; i++) {
        var b = charToBits(word[i]);
        if (b) chars.push(b);
      }
      return chars.join('000');
    });
    return wordBits.join('0000000');
  }

  function bitStringToInt32Array(bits) {
    // pad to multiple of 32
    while (bits.length % 32 !== 0) bits += '0';
    var result = [];
    for (var i = 0; i < bits.length; i += 32) {
      var chunk = bits.substring(i, i + 32);
      var num = 0;
      for (var j = 0; j < 32; j++) {
        num = (num << 1) | (chunk[j] === '1' ? 1 : 0);
      }
      result.push(num | 0); // signed 32-bit
    }
    return result;
  }

  function int32ArrayToBitString(arr) {
    var bits = '';
    for (var i = 0; i < arr.length; i++) {
      var n = arr[i];
      for (var j = 31; j >= 0; j--) {
        bits += (n >>> j) & 1 ? '1' : '0';
      }
    }
    return bits;
  }

  function bitsToMorse(bits) {
    return bits.split('0').reduce(function(acc, run) {
      if (run.length === 1) return acc + '.';
      if (run.length === 3) return acc + '-';
      return acc;
    }, '');
  }

  function decodeBitString(bits) {
    // strip trailing zeros
    bits = bits.replace(/0+$/, '');
    if (!bits) return '';

    var words = bits.split('0000000');
    return words.map(function(wordBits) {
      var chars = wordBits.split('000');
      return chars.map(function(charBits) {
        var morse = bitsToMorse(charBits);
        return reverseMorse[morse] || '';
      }).join('');
    }).join(' ');
  }

  return {
    encode: function(message) {
      var bits = messageToBitString(message);
      if (!bits) return [];
      return bitStringToInt32Array(bits);
    },
    decode: function(data) {
      if (!data || !data.length) return '';
      var bits = int32ArrayToBitString(data);
      return decodeBitString(bits);
    }
  };
})();