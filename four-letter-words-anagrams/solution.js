function anagrams(subjects, memories, players) {
  const ORDER = [
    [0, 2], [2, 1], [1, 3], [3, 0],
    [0, 3], [2, 0], [1, 2], [3, 1]
  ];

  // find starting index in ORDER
  let startIdx = ORDER.findIndex(
    ([a, b]) => a === players[0] && b === players[1]
  );

  // build memory sets for fast lookup
  const memSets = memories.map(m => new Set(m));

  // helper: sort letters to get canonical anagram key
  const sortWord = w => w.split('').sort().join('');

  // precompute anagram groups per player: for each canonical key, which words does the player know
  const playerAnagrams = memSets.map(memSet => {
    const map = new Map();
    for (const word of memSet) {
      const key = sortWord(word);
      if (!map.has(key)) map.set(key, []);
      map.get(key).push(word);
    }
    return map;
  });

  const teamScore = [0, 0];

  // which team does a player belong to
  const teamOf = p => (p === 0 || p === 1) ? 0 : 1;

  for (let i = 0; i < subjects.length; i++) {
    const roundIdx = (startIdx + i) % 8;
    const [p1, p2] = ORDER[roundIdx];
    const subject = subjects[i];
    const key = sortWord(subject);

    let p1Score = 0;
    let p2Score = 0;

    // step 1: both get 1 point if they know the subject word
    if (memSets[p1].has(subject)) p1Score += 1;
    if (memSets[p2].has(subject)) p2Score += 1;

    // get anagrams each player knows (excluding the subject word itself)
    const p1Anagrams = (playerAnagrams[p1].get(key) || []).filter(w => w !== subject);
    const p2Anagrams = (playerAnagrams[p2].get(key) || []).filter(w => w !== subject);

    // step 2: player one recalls anagrams for 2 points each
    const p1Set = new Set(p1Anagrams);
    p1Score += p1Anagrams.length * 2;

    // step 3: player two recalls anagrams NOT already recalled by player one, for 3 points each
    for (const w of p2Anagrams) {
      if (!p1Set.has(w)) {
        p2Score += 3;
      }
    }

    // round winner gets 1 final point for their team
    if (p1Score > p2Score) {
      teamScore[teamOf(p1)] += 1;
    } else if (p2Score > p1Score) {
      teamScore[teamOf(p2)] += 1;
    }
  }

  if (teamScore[0] > teamScore[1]) return 0;
  if (teamScore[1] > teamScore[0]) return 1;
  return -1;
}