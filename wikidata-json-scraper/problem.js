/*
 * Wikidata Json Scraper
 * Difficulty: 6 kyu
 *
 * Background:
 * Wikidata is a public database with over a hundred million entries in it.
 * You can find almost anything documented, from scientific articles to new
 * species scientists have found.
 *
 * To help developers use the site better, the website provides accessible
 * JSON data for every submitted item. Today, you are going to program
 * something that can read that data.
 *
 * What to do:
 * Your function should return a dictionary with the English (en) values of
 * the Identifier, Label, and Description. The dictionary should have the
 * keys "ID", "LABEL", and "DESCRIPTION", respectively.
 *
 * Be careful. Some information won't always be available in English, and
 * some might not have a value associated with any language. If the en label
 * or description aren't included, put "No Label" or "No Description".
 * Please be careful you are getting values from the "en" value, and not
 * localized english versions such as "en-uk" or en-ca.
 *
 * Formal Input/Output Requirements:
 * You are given a single argument named `url`, a string with a link to a
 * wikidata page in JSON format.
 * Your function should return a dictionary with keys "ID", "LABEL", and
 * "DESCRIPTION" containing the matching "en" values.
 */

const https = require('https');

function wikidataScraper(url) {

}