/*
 * Acronym Buster
 *
 * Laura really hates people using acronyms in her office and wants to force
 * her colleagues to remove all acronyms before emailing her.
 *
 * Any combination of three or more letters in upper case will be considered
 * an acronym. Acronyms will not be combined with lowercase letters (e.g. 'KPIs').
 * They will be kept isolated as a word/words within a string.
 *
 * Known acronyms and their replacements:
 *   KPI  -> "key performance indicators"
 *   EOD  -> "the end of the day"
 *   TBD  -> "to be decided"
 *   WAH  -> "work at home"
 *   IAM  -> "in a meeting"
 *   OOO  -> "out of office"
 *   NRN  -> "no reply necessary"
 *   CTA  -> "call to action"
 *   SWOT -> "strengths, weaknesses, opportunities and threats"
 *
 * If unknown acronyms exist, return only:
 *   '[acronym] is an acronym. I do not like acronyms. Please remove them from your email.'
 * (first unknown acronym only)
 *
 * If all acronyms are known, return the altered string with proper sentence capitalization.
 */
function acronymBuster(message) {
}