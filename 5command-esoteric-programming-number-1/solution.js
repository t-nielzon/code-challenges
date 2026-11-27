class Output {
  constructor(out, debug) {
    this.output = out;
    this.debug = debug;
  }
}

class FiveCommands {
  read(input, debug) {
    const tape = {}; // sparse array using object for unlimited tape
    let pointer = 0;
    let output = '';
    
    for (let cmd of input) {
      switch(cmd) {
        case '+':
          pointer++;
          break;
        case '-':
          pointer--;
          break;
        case '^':
          tape[pointer] = (tape[pointer] || 0) + 1;
          break;
        case 'v':
          tape[pointer] = (tape[pointer] || 0) - 1;
          break;
        case '*':
          const val = tape[pointer] || 0;
          output += val;
          break;
      }
    }
    
    let debugOutput = [];
    if (debug) {
      const keys = Object.keys(tape).sort((a, b) => a - b);
      debugOutput = keys.map(k => `${k}->${tape[k]}`);
    }
    
    return new Output(output, debugOutput);
  }
}