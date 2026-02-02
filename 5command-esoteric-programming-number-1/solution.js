class Output {
  constructor(out, debug) {
    this.output = out;
    this.debug = debug;
  }
}

class FiveCommands {
  read(input, debug) {
    let tape = [0];
    let pointer = 0;
    let output = '';

    for (const cmd of input) {
      switch (cmd) {
        case '+':
          pointer++;
          if (pointer >= tape.length) {
            tape.push(0);
          }
          break;
        case '-':
          if (pointer === 0) {
            tape.unshift(0);
          } else {
            pointer--;
          }
          break;
        case '^':
          tape[pointer]++;
          break;
        case 'v':
          tape[pointer]--;
          break;
        case '*':
          output += tape[pointer].toString();
          break;
      }
    }

    let debugArray = [];
    if (debug) {
      for (let i = 0; i < tape.length; i++) {
        debugArray.push(`${i}->${tape[i]}`);
      }
    }

    return new Output(output, debugArray);
  }
}