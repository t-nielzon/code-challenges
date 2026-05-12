function stat(strg) {
  if (strg === "") return "";

  const times = strg.split(/,\s*/).map(part => {
    const [h, m, s] = part.split("|").map(Number);
    return h * 3600 + m * 60 + s;
  });

  const sorted = [...times].sort((a, b) => a - b);
  const n = sorted.length;

  const range = sorted[n - 1] - sorted[0];
  const average = Math.trunc(sorted.reduce((a, b) => a + b, 0) / n);
  const median = n % 2 === 1
    ? sorted[(n - 1) / 2]
    : Math.trunc((sorted[n / 2 - 1] + sorted[n / 2]) / 2);

  const fmt = (sec) => {
    const h = Math.floor(sec / 3600);
    const m = Math.floor((sec % 3600) / 60);
    const s = sec % 60;
    const pad = (n) => String(n).padStart(2, "0");
    return `${pad(h)}|${pad(m)}|${pad(s)}`;
  };

  return `Range: ${fmt(range)} Average: ${fmt(average)} Median: ${fmt(median)}`;
}