function format(str) {
  // Check for headers (1-6 hashtags followed by space)
  const headerMatch = str.match(/^(#+)\s+(.+)$/);
  if (headerMatch) {
    const level = Math.min(headerMatch[1].length, 6);
    const content = applyStrong(headerMatch[2]);
    return `< h${level}>${content}< /h${level}>`;
  }
  
  // Check for list items (asterisk followed by space)
  const listMatch = str.match(/^\*\s+(.+)$/);
  if (listMatch) {
    const content = applyStrong(listMatch[1]);
    return `< li>${content}< /li>`;
  }
  
  // Default to paragraph
  const content = applyStrong(str);
  return `< p>${content}< /p>`;
}

function applyStrong(text) {
  // Replace **text** with < strong>text< /strong> (non-greedy matching)
  return text.replace(/\*\*(.+?)\*\*/g, '< strong>$1< /strong>');
}