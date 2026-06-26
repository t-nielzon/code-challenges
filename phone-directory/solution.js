function phone(strng, num) {
  const lines = strng.split("\n").filter(l => l.includes("+" + num));

  if (lines.length === 0) return `Error => Not found: ${num}`;
  if (lines.length > 1) return `Error => Too many people: ${num}`;

  let line = lines[0];

  // extract name between < and >
  const name = line.match(/<(.*?)>/)[1];

  // remove the name (with brackets) and the phone number from the line
  line = line.replace(/<.*?>/, " ").replace("+" + num, " ");

  // whatever non phone/name content remains is the address; tidy up punctuation
  const address = line
    .replace(/[^A-Za-z0-9\s.\-_]/g, " ")
    .replace(/_/g, " ")
    .replace(/\s+/g, " ")
    .trim();

  return `Phone => ${num}, Name => ${name}, Address => ${address}`;
}