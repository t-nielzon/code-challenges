/**
 * Translate DNA in 6 frames
 * 
 * In genetics, a reading frame divides a DNA sequence into consecutive 
 * non-overlapping triplets (codons). A single strand has 3 forward frames,
 * and the reverse complement strand provides 3 more reverse frames.
 * 
 * This function translates DNA on all 6 frames using the standard genetic code.
 * 
 * @param {string} dna - The DNA sequence
 * @param {number[]} frames - Array of frame numbers (1, 2, 3 for forward; -1, -2, -3 for reverse)
 * @returns {string[]} - Array of translated amino acid sequences
 */
function translateDNA(dna, frames = [1, 2, 3, -1, -2, -3]) {
  
}