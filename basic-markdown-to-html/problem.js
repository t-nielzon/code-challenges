/*
You've been given some texts for your website which you need to turn into HTML.
Unfortunately for you, the texts also contain markdown, which you need to take account of too.

In this Kata you'll write a function "format()" which transforms markdown elements to HTML:
- Bullet points (list items): * followed by space
- Headers (levels 1 through 6): # at beginning followed by space (max 6 hashtags for header level)
- Bold emphasis: text surrounded by double asterisks using non-greedy matching

Input is a single line of Markdown with no newlines.

Rules:
1. Strong emphasis: **text** becomes < strong>text< /strong> (non-greedy matching)
2. Headers: # text becomes < h1>text< /h1> (1-6 hashtags; 7+ treated as h6 with extras as content)
3. List items: * text becomes < li>text< /li>
4. Default: wrap in < p>text< /p>
5. No extra spaces between tags and content
*/

function format(str) {
  // Implementation here
}