function translateDNA(dna, frames = [1, 2, 3, -1, -2, -3]) {
  const codons = createCodonTable();
  const results = [];
  
  for (const frame of frames) {
    let sequence, start;
    
    if (frame > 0) {
      sequence = dna;
      start = frame - 1;
    } else {
      sequence = getReverseComplement(dna);
      start = -frame - 1;
    }
    
    let translation = '';
    for (let i = start; i + 3 <= sequence.length; i += 3) {
      const codon = sequence.substring(i, i + 3);
      translation += codons[codon];
    }
    
    results.push(translation);
  }
  
  return results;
}

function getReverseComplement(dna) {
  const complement = { A: 'T', T: 'A', G: 'C', C: 'G' };
  let result = '';
  for (let i = dna.length - 1; i >= 0; i--) {
    result += complement[dna[i]];
  }
  return result;
}

function createCodonTable() {
  const aas = "FFLLSSSSYY**CC*WLLLLPPPPHHQQRRRRIIIMTTTTNNKKSSRRVVVVAAAADDEEGGGG";
  const base1 = "TTTTTTTTTTTTTTTTCCCCCCCCCCCCCCCCAAAAAAAAAAAAAAAAGGGGGGGGGGGGGGGG";
  const base2 = "TTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGG";
  const base3 = "TCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAG";
  
  const codons = {};
  for (let i = 0; i < aas.length; i++) {
    codons[base1[i] + base2[i] + base3[i]] = aas[i];
  }
  
  return codons;
}