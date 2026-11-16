function format(text) {
  // Check for headers (1 or more # followed by space, cap at 6 levels)
  const headerMatch = text.match(/^(#+) /);
  if (headerMatch) {
    const level = Math.min(headerMatch[1].length, 6);
    const content = text.slice(headerMatch[0].length);
    const processedContent = processBold(content);
    return `<h${level}>${processedContent}</h${level}>`;
  }
  
  // Check for list items (* followed by space)
  if (text.startsWith('* ')) {
    const content = text.slice(2);
    const processedContent = processBold(content);
    return `<li>${processedContent}</li>`;
  }
  
  // Default: paragraph
  const processedContent = processBold(text);
  return `<p>${processedContent}</p>`;
}

function processBold(text) {
  return text.replace(/\*\*(.+?)\*\*/g, '<strong>$1</strong>');
}