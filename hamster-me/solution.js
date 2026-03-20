function hamsterMe(code, message) {
  const codeLetters = [...new Set(code.split(''))].sort();
  
  // assign each letter of the alphabet a column header and row number
  const assigned = {};
  
  // code letters get row 1
  for (const ch of codeLetters) {
    assigned[ch] = { base: ch, row: 1 };
  }
  
  // for each code letter, walk forward through the alphabet assigning rows
  for (const ch of codeLetters) {
    let charCode = ch.charCodeAt(0);
    let row = 2;
    while (true) {
      charCode++;
      if (charCode > 122) charCode = 97; // wrap z -> a
      const letter = String.fromCharCode(charCode);
      if (assigned[letter]) break; // hit another column's territory
      assigned[letter] = { base: ch, row };
      row++;
    }
  }
  
  // handle unassigned letters: they go after the last code letter alphabetically
  const lastCode = codeLetters[codeLetters.length - 1];
  let row = 2;
  for (let i = 0; i < 26; i++) {
    const letter = String.fromCharCode(((lastCode.charCodeAt(0) - 97 + i) % 26) + 97);
    if (!assigned[letter]) {
      assigned[letter] = { base: lastCode, row };
      row++;
    } else if (letter !== lastCode) {
      // already assigned, but we need to keep counting from lastCode
      // actually the loop above handles the gap case; if assigned, skip
    }
  }
  
  // re-approach: the column assignment via "walk forward until hitting another code letter"
  // may leave gaps when code letters aren't contiguous. Let me reconsider.
  // Actually the walk does handle it: for each code letter we walk forward until we hit
  // an already-assigned letter. Letters not reachable this way (before the first code letter
  // wrapping around) get appended after the last code letter.
  // But the append logic above is wrong. Let me fix it.
  
  // Check for any unassigned letters and append them after lastCode's column
  const lastCodeAssigned = Object.entries(assigned)
    .filter(([, v]) => v.base === lastCode)
    .map(([, v]) => v.row);
  let maxRow = Math.max(...lastCodeAssigned);
  
  // walk from lastCode+1 around the alphabet, assign any unassigned
  for (let i = 1; i <= 26; i++) {
    const letter = String.fromCharCode(((lastCode.charCodeAt(0) - 97 + i) % 26) + 97);
    if (!assigned[letter]) {
      maxRow++;
      assigned[letter] = { base: lastCode, row: maxRow };
    }
  }
  
  return message.split('').map(ch => assigned[ch].base + assigned[ch].row).join('');
}