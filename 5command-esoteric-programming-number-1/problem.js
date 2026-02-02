/*
 * 5command - Esoteric programming #1
 *
 * 5command is an Esoteric programming language which has 5 commands.
 * Your task is to create an interpreter for a 5command input.
 *
 * Commands:
 * + : move the pointer right on the tape
 * - : move the pointer left on the tape
 * ^ : increment the value on the tape where the pointer is currently
 * v : decrement the value on the tape where the pointer is currently
 * * : print to tape the number at the pointer (without a space)
 *
 * The tape used is unlimited - expand at the start if moving to -1.
 *
 * Input: read(input, debug)
 * - input: string of commands (e.g., "^*")
 * - debug: boolean determining if debug array is returned
 *
 * Output: Output object with:
 * - output: result as a string
 * - debug: array representation of the tape (or empty array if debug is false)
 *   Format: ['index->value', ...]
 */

class Output {
  constructor(out, debug) {
    this.output = out;
    this.debug = debug;
  }
}

class FiveCommands {
  read(input, debug) {
    // Your implementation here
  }
}