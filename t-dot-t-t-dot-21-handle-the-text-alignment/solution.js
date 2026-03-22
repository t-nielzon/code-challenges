function formatScoreList(list) {
  const avg = list.reduce((s, x) => s + x.score, 0) / list.length;
  const avgStr = avg.toFixed(1);

  const nameLabel = "Name";
  const scoreLabel = "Score";
  const avgName = "The average score";

  const allNames = list.map(x => x.name).concat(avgName);
  const leftW = Math.max(nameLabel.length, ...allNames.map(n => n.length));

  const allScores = list.map(x => String(x.score) + " pts").concat(avgStr + " pts");
  const rightW = Math.max(scoreLabel.length, ...allScores.map(s => s.length));

  const sep = `+${"-".repeat(leftW + 1)}+${"-".repeat(rightW + 1)}+`;

  // center name header: extra space goes to right
  const totalPad = leftW - nameLabel.length;
  const padL = Math.floor(totalPad / 2);
  const padR = totalPad - padL;
  const headerLeft = " ".repeat(padL) + nameLabel + " ".repeat(padR);

  // center score header
  const totalPadS = rightW - scoreLabel.length;
  const padSL = Math.floor(totalPadS / 2);
  const padSR = totalPadS - padSL;
  const headerRight = " ".repeat(padSL) + scoreLabel + " ".repeat(padSR);

  // header row has padding on both sides (space before | and after |)
  const header = `|${" " + headerLeft}|${" " + headerRight}|`;

  const rows = [sep, header, sep];

  const formatRow = (name, scoreStr) => {
    const left = name + " ".repeat(leftW - name.length);
    const right = scoreStr + " ".repeat(rightW - scoreStr.length);
    return `|${left + " "}|${right}|`;
  };

  // wait — looking at the examples more carefully:
  // |Tom              |100  pts|
  // left width includes the padding to the right, no extra space before |
  // right width: no extra space
  // but header: |      Name       | Score  |
  // header left has +1 width (space on each side of content area)
  // actually the separator is +---...---+---...---+
  // let me re-examine

  // separator: +-----------------+--------+
  // that's 17 dashes and 8 dashes
  // "The average score" = 17 chars, so leftW = 17
  // "100  pts" = 8 chars? no, "100  pts" has 2 spaces = 8 chars? 1+0+0+space+space+p+t+s = 8
  // but rightW from "Score" with padding = " Score  " ... hmm

  // Let me reconsider. The inner widths (between pipes) are:
  // Left column inner width = max(all left content lengths)
  // Right column inner width = max(all right content lengths)
  // Header adds spaces for centering within that width
  // But header row: |      Name       | Score  |
  // "      Name       " = 17+1? no...

  // +-----------------+--------+  => 17 dashes, 8 dashes
  // |      Name       | Score  |  => 17 chars inside, 8 chars inside? 
  // "      Name       " = 6+4+7 = 17. Yes.
  // " Score  " = 1+5+2 = 8. Yes.
  // |Tom              |100  pts| => "Tom              " = 17, "100  pts" = 8

  // So leftW = 17 = len("The average score"), rightW = 8
  // But "100  pts" = "100" + "  " + "pts" = 8? 3+2+3=8. And rightW should be max of all score strings.
  // "Score" = 5, needs to fit in 8 with centering: " Score  " (1 left, 2 right) — extra right.
  // Wait: " Score  " has space+Score+2spaces = 8. Left pad=1, right pad=2. Total pad=3, floor(3/2)=1 left, 2 right. Yes.

  // For data rows: left is flush-left padded to leftW, right is "both ends aligned" to rightW
  // "both ends aligned" for scores means: score number flush left, "pts" flush right
  // "100  pts" in width 8: "100" then spaces then "pts", total 8
  // "95   pts" in width 8: "95" then spaces then "pts", total 8

  // So the right column format is: scoreNum + spaces + "pts", total = rightW

  // Re-derive rightW: it must fit " Score " centered AND "X...pts" for all scores
  // Actually from example 2: score col = 7 dashes
  // "1   pts" = 7, "2.5 pts" = 7, " Score " = 7 (1+5+1)
  // So rightW = max(len("X pts") for all scores including avg, also at least scoreLabel.length + 2 for centering?)

  // Nope. Let me just compute: the score strings are like "100  pts", "95   pts" etc.
  // Format: `${scoreNum} pts` but padded so total = rightW with "pts" at the right end.
  // Actually it's simpler: rightW = max(scoreLabel + 2, maxScoreLen + 1 + 3) where +2 is for " Score " min padding
  
  // Let me just recalculate properly from examples.
  // Example 1: scores 100,95,90. avg=95.0
  // Score strings: "100", "95", "90", "95.0" => longest score num = 3 ("100") or 4 ("95.0")
  // "95.0 pts" = 8, "100  pts" = 8... so both are 8
  // rightW = 8 = max(len(s) for s in formatted score strings)
  // where formatted = scoreNum padded to maxScoreNumLen + " pts"? 
  // No: "100  pts" has 2 spaces, "95.0 pts" has 1 space
  // Actually: content = scoreNum + padding + "pts", total = rightW
  // So rightW must be >= len(scoreNum) + 1 + 3 for all scoreNums, and >= scoreLabel + 2

  // Hmm let me just look at it as: rightW = max of all right-side content widths
  // Right content for scores: we need scoreNum left-aligned and "pts" right-aligned within rightW
  // So rightW >= max(len(scoreNum) + 1 + 3) for all score nums (need at least 1 space between num and pts)
  // And rightW >= len(" " + scoreLabel + " ") for centering? Actually the header centering adapts to rightW.

  // From example 2: scores 1,2,3,4, avg 2.5
  // rightW = 7: " Score " (centered), "1   pts" "2.5 pts"
  // max score num len = 3 ("2.5"), so 3+1+3=7. And "Score"=5, centered in 7 = " Score " (1+5+1). Works.
  
  // Example 1: scores 100,95,90, avg 95.0
  // max score num len = 4 ("95.0"), 4+1+3=8. "Score"=5 centered in 8 = " Score  " (1+5+2). Works.
  
  // So: rightW = max(maxScoreNumWidth + 1 + 3, scoreLabel.length + 2)
  // where scoreLabel.length + 2 ensures at least 1 space on each side for centering

  // Actually let me check: what if scoreLabel is longer? Like if all scores are single digit.
  // scores all 1 digit: maxScoreNumWidth=1, 1+1+3=5. "Score"=5, 5+2=7. rightW=7. " Score " centered in 7. Yes that matches example 2.
  
  // Wait example 2 has avg "2.5" which is 3 chars, so maxScoreNumWidth=3, 3+1+3=7. So rightW=7 from scores alone.
  // Let me try: all scores single digit, avg also single digit like scores=[{score:5},{score:5}] avg=5.0 (3 chars)
  // maxScoreNumWidth = 3 ("5.0"), rightW = max(7, 7) = 7. OK.
  
  // What about scores=[{score:5}], avg=5.0? Same thing.
  
  // I think the formula is just: rightW = max(maxScoreNumWidth + 4, scoreLabel.length + 2)
  // where 4 = 1 space + "pts"

  const scoreNums = list.map(x => String(x.score)).concat(avgStr);
  const maxScoreNumW = Math.max(...scoreNums.map(s => s.length));
  const rightWidth = Math.max(maxScoreNumW + 4, scoreLabel.length + 2);
  const leftWidth = Math.max(...allNames.map(n => n.length));

  const separator = `+${"-".repeat(leftWidth + 1)}+${"-".repeat(rightWidth + 1)}+`;

  // center header
  const cen = (text, w) => {
    const t = w - text.length;
    const l = Math.floor(t / 2);
    const r = t - l;
    return " ".repeat(l) + text + " ".repeat(r);
  };

  const hdrLeft = cen(nameLabel, leftWidth);
  const hdrRight = cen(scoreLabel, rightWidth);
  const headerRow = `|${hdrLeft} | ${hdrRight}|`;

  // wait, looking at example output more carefully:
  // +-----------------+--------+
  // |      Name       | Score  |
  // The separator has leftWidth+1 dashes (17 chars name + 1 padding?) 
  // Hmm: "+-----------------+--------+" 
  // Count: 17 dashes, 8 dashes
  // "|      Name       | Score  |"
  // Between first two pipes: "      Name       " = 17+1 = 18? No...
  // Let me count character by character:
  // |      Name       | Score  |
  // Pipe, 6 spaces, "Name", 7 spaces, pipe, space, "Score", 2 spaces, pipe
  // Inner left: 6+4+7 = 17. Inner right: 1+5+2 = 8.
  // But separator has 17+1 dashes? No: +-----------------+--------+
  // That's +, 17 dashes, +, 8 dashes, +. Hmm wait:
  // "+-----------------+--------+"
  // Let me count the dashes: 17 and... 8? Let me recount.
  // After first +: "-----------------" that's 17.
  // After second +: "--------" that's 8.
  // So separator inner widths are 17 and 8.
  
  // But header inner widths are also 17 and 8:
  // "|      Name       | Score  |" 
  // "      Name       " = 17 chars (6+4+7)
  // " Score  " = 8 chars (1+5+2)
  // So no extra padding between separator and content. Good.

  // Data rows: "|Tom              |100  pts|"
  // "Tom              " = 17 chars (3+14)
  // "100  pts" = 8 chars
  // Same widths. 

  // So separator = "+" + "-".repeat(leftWidth) + "+" + "-".repeat(rightWidth) + "+"
  // Hmm but wait, leftWidth=17 for the name column. Is it 17 or 18?
  // "The average score" = 17 chars. And the cell content is "|The average score|"
  // The inner width between pipes is exactly 17. So leftWidth = 17. But earlier I said separator had 17 dashes too, and the header row also has 17 chars between pipes. So everything is consistent with leftWidth = 17 and rightWidth = 8, with NO extra padding.

  // But look at the header: "| Score  |" — there's a space before "Score". And "|      Name       |" has no leading space before the centered content.
  // Actually: "|      Name       | Score  |"
  // The left cell is "      Name       " (17 chars). 
  // The right cell is " Score  " (8 chars).
  // In the centered calculation for right: total pad = 8-5=3, floor(3/2)=1 left, 2 right. " Score  ". 

  // For the separator: "+-...17...-+-...8...-+" — but wait:
  // +-----------------+--------+
  //                    ^^^^^^^^
  // That's 8 dashes. So separator = "+" + "-".repeat(17) + "+" + "-".repeat(8) + "+"
  // Which means no extra +1 anywhere. The inner widths ARE the column widths.

  // But I see " Score  " has a leading space before Score — that's part of the centering, not extra padding.
  // Hmm wait, I'm second-guessing myself. Let me look at example 2:
  // +-----------------+-------+
  // |      Name       | Score |
  // Here rightWidth should be 7 ("1   pts" = 7 chars, or "2.5 pts" = 7 chars).
  // " Score " centered in 7: 7-5=2, floor(2/2)=1 left, 1 right. " Score " = 7. Yes.
  // Separator: 17 dashes and 7 dashes. 

  // OK so my analysis is correct. Let me also check data rows in example 2:
  // |Tom              |1   pts|
  // Left: "Tom              " = 17 chars? "Tom" + 14 spaces = 17. Yes.
  // Right: "1   pts" = 7. "1" + 3 spaces + "pts" = 7. Yes.

  // So the format for data rows: left cell = name.padEnd(leftWidth), right cell = scoreNum + spaces + "pts" totaling rightWidth
  // And the row is "|" + leftCell + "|" + rightCell + "|"

  // BUT WAIT. Look more carefully at example 1 output:
  // +-----------------+--------+
  // |      Name       | Score  |
  // The first separator line... Let me count again very carefully.
  // "+-----------------+--------+"
  // +, then 17 dashes, +, then 8 dashes, +. Total = 1+17+1+8+1 = 28 chars.
  // "|      Name       | Score  |"
  // |, "      Name       ", |, " Score  ", |. Inner: 17 + 8 = 25, plus 3 pipes = 28. Consistent!

  // "|Tom              |100  pts|"
  // |, "Tom              " (17), |, "100  pts" (8), |. Total = 28. Good.

  // So NO extra spaces/padding beyond the column widths. Perfect.

  // Now let me just verify my rightWidth formula with example 1:
  // Scores: 100, 95, 90. Avg: 95.0. Score nums: "100","95","90","95.0". maxScoreNumW = 4 ("95.0").
  // rightWidth = max(4+4, 5+2) = max(8,7) = 8. Correct!

  // Example 2: Scores: 1,2,3,4. Avg: 2.5. Score nums: "1","2","3","4","2.5". maxScoreNumW = 3.
  // rightWidth = max(3+4, 5+2) = max(7,7) = 7. Correct!

  const sep2 = `+${"-".repeat(leftWidth)}+${"-".repeat(rightWidth)}+`;

  const centerIn = (text, w) => {
    const pad = w - text.length;
    const l = Math.floor(pad / 2);
    const r = pad - l;
    return " ".repeat(l) + text + " ".repeat(r);
  };

  const hLeft = centerIn(nameLabel, leftWidth);
  const hRight = centerIn(scoreLabel, rightWidth);
  const hRow = `|${hLeft}|${hRight}|`;

  const makeScoreCell = (numStr) => {
    const ptsStr = "pts";
    const padding = rightWidth - numStr.length - ptsStr.length;
    return numStr + " ".repeat(padding) + ptsStr;
  };

  const makeRow = (name, scoreNumStr) => {
    const left = name.padEnd(leftWidth);
    const right = makeScoreCell(scoreNumStr);
    return `|${left}|${right}|`;
  };

  const lines = [sep2, hRow, sep2];
  for (const item of list) {
    lines.push(makeRow(item.name, String(item.score)));
    lines.push(sep2);
  }
  lines.push(makeRow(avgName, avgStr));
  lines.push(sep2);

  return lines.join("\n");
}