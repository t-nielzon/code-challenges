/*
 * 7-segment converter
 * A 7-segment display converter that turns digits 0-9 into binary instructions.
 * 
 * The display has 7 segments arranged like:
 *  _ 
 * |_|
 * | |
 *  ¯ 
 * 
 * Each segment corresponds to a bit position (reading right to left):
 * - Bit 0: Top Horizontal
 * - Bit 1: Middle Horizontal
 * - Bit 2: Bottom Horizontal
 * - Bit 3: Top-Left Vertical
 * - Bit 4: Top-Right Vertical
 * - Bit 5: Bottom-Left Vertical
 * - Bit 6: Bottom-Right Vertical
 * 
 * The function returns the decimal representation of the binary instruction
 * for each digit 0-9.
 */

function converter(num) {
  
}