function chordChecker(chord) {
  const notes = chord.split(' ');
  
  if (notes.length !== 3) {
    return 'Not a chord';
  }
  
  const chromatic = ['C', ['C#', 'Db'], 'D', ['D#', 'Eb'], 'E', 'F', ['F#', 'Gb'], 'G', ['G#', 'Ab'], 'A', ['A#', 'Bb'], 'B'];
  
  const noteToPosition = {};
  for (let i = 0; i < chromatic.length; i++) {
    const note = chromatic[i];
    if (Array.isArray(note)) {
      note.forEach(n => {
        noteToPosition[n] = i;
      });
    } else {
      noteToPosition[note] = i;
    }
  }
  
  const positions = notes.map(note => noteToPosition[note]);
  
  if (positions.some(pos => pos === undefined)) {
    return 'Not a chord';
  }
  
  const interval1 = (positions[1] - positions[0] + 12) % 12;
  const interval2 = (positions[2] - positions[1] + 12) % 12;
  const interval3 = (positions[2] - positions[0] + 12) % 12;
  
  if (interval3 !== 7) {
    return 'Not a chord';
  }
  
  if (interval1 === 3 && interval2 === 4) {
    return 'Minor';
  } else if (interval1 === 4 && interval2 === 3) {
    return 'Major';
  } else {
    return 'Not a chord';
  }
}