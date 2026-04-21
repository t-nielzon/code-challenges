function rank(st, we, n) {
  if (!st) return "No participants";
  const names = st.split(",");
  if (n > names.length) return "Not enough participants";

  const scored = names.map((name, i) => {
    const sum = name
      .toLowerCase()
      .split("")
      .reduce((acc, ch) => acc + (ch.charCodeAt(0) - 96), 0);
    const som = name.length + sum;
    return { name, winning: som * we[i] };
  });

  scored.sort((a, b) => {
    if (b.winning !== a.winning) return b.winning - a.winning;
    return a.name < b.name ? -1 : a.name > b.name ? 1 : 0;
  });

  return scored[n - 1].name;
}