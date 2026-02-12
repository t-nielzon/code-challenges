function anagrams(subjects, memories, players) {
  const order = [
    [0, 2], [2, 1], [1, 3], [3, 0],
    [0, 3], [2, 0], [1, 2], [3, 1]
  ];

  // find starting index in the order
  let orderIdx = order.findIndex(
    ([a, b]) => a === players[0] && b === players[1]
  );

  // build sets for each player's memory, keyed by sorted letters
  const memorySets = memories.map(mem => new Set(mem));

  // helper: sort letters of a word to find anagram key
  const sortWord = w => w.split('').sort().join('');

  // build a map from sorted-key to all words in all memories
  const allWords = new Set();
  for (const mem of memories) for (const w of mem) allWords.add(w);
  for (const w of subjects) allWords.add(w);

  const anagramMap = new Map();
  for (const w of allWords) {
    const key = sortWord(w);
    if (!anagramMap.has(key)) anagramMap.set(key, []);
    anagramMap.get(key).push(w);
  }

  const teamScore = [0, 0];
  const getTeam = p => (p === 0 || p === 1) ? 0 : 1;

  for (let i = 0; i < subjects.length; i++) {
    const [p1, p2] = order[orderIdx];
    orderIdx = (orderIdx + 1) % 8;

    const subject = subjects[i];
    const key = sortWord(subject);
    const allAnagrams = anagramMap.has(key) ? anagramMap.get(key) : [];

    let score1 = 0;
    let score2 = 0;

    // 1 point each if they know the subject word
    const p1KnowsSubject = memorySets[p1].has(subject);
    const p2KnowsSubject = memorySets[p2].has(subject);
    if (p1KnowsSubject) score1 += 1;
    if (p2KnowsSubject) score2 += 1;

    // find anagrams (excluding the subject word itself)
    const p1Anagrams = [];
    const p1AnagramSet = new Set();
    for (const w of allAnagrams) {
      if (w === subject) continue;
      if (memorySets[p1].has(w)) {
        p1Anagrams.push(w);
        p1AnagramSet.add(w);
      }
    }
    // p1 gets 2 points per anagram
    score1 += p1Anagrams.length * 2;

    // p2 gets 3 points for anagrams p1 didn't know
    for (const w of allAnagrams) {
      if (w === subject) continue;
      if (memorySets[p2].has(w) && !p1AnagramSet.has(w)) {
        score2 += 3;
      }
    }

    if (score1 > score2) {
      teamScore[getTeam(p1)] += 1;
    } else if (score2 > score1) {
      teamScore[getTeam(p2)] += 1;
    }
  }

  if (teamScore[0] > teamScore[1]) return 0;
  if (teamScore[1] > teamScore[0]) return 1;
  return -1;
}

module.exports = { anagrams };