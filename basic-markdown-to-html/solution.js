function format(text) {
  function processStrong(str) {
    return str.replace(/\*\*(.+?)\*\*/g, '< strong>$1< /strong>');
  }
  
  const headerMatch = text.match(/^(#+) (.*)$/);
  if (headerMatch) {
    const level = Math.min(headerMatch[1].length, 6);
    const content = processStrong(headerMatch[2]);
    return `< h${level}>${content}< /h${level}>`;
  }
  
  if (text.match(/^\* /)) {
    const content = processStrong(text.substring(2));
    return `< li>${content}< /li>`;
  }
  
  const content = processStrong(text);
  return `< p>${content}< /p>`;
}