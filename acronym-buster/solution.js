function acronymBuster(message) {
  const known = {
    KPI: "key performance indicators",
    EOD: "the end of the day",
    TBD: "to be decided",
    WAH: "work at home",
    IAM: "in a meeting",
    OOO: "out of office",
    NRN: "no reply necessary",
    CTA: "call to action",
    SWOT: "strengths, weaknesses, opportunities and threats",
  };

  const acronymRegex = /\b[A-Z]{3,}\b/g;
  const matches = message.match(acronymRegex);

  if (matches) {
    const unknown = matches.find((m) => !known[m]);
    if (unknown) {
      return `${unknown} is an acronym. I do not like acronyms. Please remove them from your email.`;
    }
  }

  let result = message.replace(acronymRegex, (match) => known[match]);

  // capitalize first letter of each sentence
  result = result.replace(/(^|[.]\s*)([a-z])/g, (_, pre, letter) =>
    pre + letter.toUpperCase()
  );

  return result;
}