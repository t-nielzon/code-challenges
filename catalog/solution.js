function catalog(s, article) {
  const re = /<prod><name>([^<]+)<\/name><prx>([^<]+)<\/prx><qty>([^<]+)<\/qty><\/prod>/g;
  const results = [];
  const wordRe = new RegExp('\\b' + article.replace(/[.*+?^${}()|[\]\\]/g, '\\$&') + '\\b');
  let m;
  while ((m = re.exec(s)) !== null) {
    const [, name, prx, qty] = m;
    if (wordRe.test(name)) {
      results.push(`${name} > prx: $${prx} qty: ${qty}`);
    }
  }
  return results.length ? results.join('\n') : 'Nothing';
}