/*
5command is an Esoteric programming language which has 5 commands.
Your task is to create an interpreter for a 5command input.

Commands:
+ move the pointer right on the tape
- move the pointer left on the tape
^ increment the value on the tape where the pointer is currently
v decrement the value on the tape where the pointer is currently
* print to tape the number at the pointer (without a space)

The tape is unlimited and extends infinitely in both directions.
Each cell starts at 0.

Returns an Output object with:
- output: string of printed values
- debug: array of "position->value" for modified cells (empty if debug is false)
*/

class Output {
  constructor(out, debug) {
    this.output = out;
    this.debug = debug;
  }
}

class FiveCommands {
  read(input, debug) {
    // TODO: implement interpreter
  }
}