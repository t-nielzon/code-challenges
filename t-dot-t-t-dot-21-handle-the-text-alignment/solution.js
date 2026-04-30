function formatScoreList(list) {
  const avg = list.reduce((s, x) => s + x.score, 0) / list.length;
  const avgStr = avg.toFixed(1);

  const leftStrings = list.map(x => x.name).concat(["The average score", "Name"]);
  const leftWidth = Math.max(...leftStrings.map(s => s.length));

  const scoreStrs = list.map(x => String(x.score)).concat([avgStr]);
  const maxScoreLen = Math.max(...scoreStrs.map(s => s.length));
  const rightWidth = maxScoreLen + 4;

  const sep = "+" + "-".repeat(leftWidth) + "+" + "-".repeat(rightWidth) + "+";

  const center = (text, width) => {
    const total = width - text.length;
    const left = Math.floor(total / 2);
    const right = total - left;
    return " ".repeat(left) + text + " ".repeat(right);
  };

  const formatRight = (s) => s.padEnd(maxScoreLen) + " pts";

  const lines = [sep];
  lines.push("|" + center("Name", leftWidth) + "|" + center("Score", rightWidth) + "|");
  lines.push(sep);

  for (const item of list) {
    lines.push("|" + item.name.padEnd(leftWidth) + "|" + formatRight(String(item.score)) + "|");
    lines.push(sep);
  }

  lines.push("|" + "The average score".padEnd(leftWidth) + "|" + formatRight(avgStr) + "|");
  lines.push(sep);

  return lines.join("\n");
}