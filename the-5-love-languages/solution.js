function findLoveLanguage(partner, weeks) {
  const LOVE_LANGUAGES = ["words", "acts", "gifts", "time", "touch"];
  const totalTries = weeks * 7;
  const triesPerLanguage = Math.floor(totalTries / LOVE_LANGUAGES.length);
  
  const scores = {};
  LOVE_LANGUAGES.forEach(lang => {
    scores[lang] = 0;
  });
  
  // Test each language multiple times and count positive responses
  for (let lang of LOVE_LANGUAGES) {
    for (let i = 0; i < triesPerLanguage; i++) {
      if (partner.response(lang) === "positive") {
        scores[lang]++;
      }
    }
  }
  
  // Return the language with the highest score
  return LOVE_LANGUAGES.reduce((best, lang) => 
    scores[lang] > scores[best] ? lang : best
  );
}