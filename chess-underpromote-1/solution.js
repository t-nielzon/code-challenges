function underpromote(start, target) {
  if (start === target) return 0;

  const parse = (s) => [s.charCodeAt(0) - 96, parseInt(s[1])];
  const [sf, sr] = parse(start);
  const [tf, tr] = parse(target);

  // bfs over states: [file, rank, piece] where piece: 0=pawn, 1=queen, 2=knight
  const key = (f, r, p) => f * 100 + r * 10 + p;
  const visited = new Set([key(sf, sr, 0)]);
  const queue = [[sf, sr, 0, 0]];

  while (queue.length) {
    const [f, r, piece, moves] = queue.shift();
    const nextStates = [];

    if (piece === 0) {
      const nr = r + 1;
      if (nr <= 8) {
        if (nr === 8) {
          // pawn arrives at 8th rank: can stay as pawn (only useful if target), or promote
          nextStates.push([f, nr, 0], [f, nr, 1], [f, nr, 2]);
        } else {
          nextStates.push([f, nr, 0]);
        }
      }
    } else if (piece === 1) {
      // queen: slides in 8 directions
      for (const [df, dr] of [[-1,-1],[-1,0],[-1,1],[0,-1],[0,1],[1,-1],[1,0],[1,1]]) {
        let nf = f + df, nr = r + dr;
        while (nf >= 1 && nf <= 8 && nr >= 1 && nr <= 8) {
          nextStates.push([nf, nr, 1]);
          nf += df;
          nr += dr;
        }
      }
    } else {
      // knight: 8 L-shaped moves
      for (const [df, dr] of [[-2,-1],[-2,1],[-1,-2],[-1,2],[1,-2],[1,2],[2,-1],[2,1]]) {
        const nf = f + df, nr = r + dr;
        if (nf >= 1 && nf <= 8 && nr >= 1 && nr <= 8) {
          nextStates.push([nf, nr, 2]);
        }
      }
    }

    for (const [nf, nr, np] of nextStates) {
      if (nf === tf && nr === tr) return moves + 1;
      const k = key(nf, nr, np);
      if (!visited.has(k)) {
        visited.add(k);
        queue.push([nf, nr, np, moves + 1]);
      }
    }
  }
}