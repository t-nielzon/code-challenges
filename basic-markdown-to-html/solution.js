function format(str) {
  let content = str;
  let tag = 'p';
  
  // Check for header
  const headerMatch = str.match(/^(#{1,})\s/);
  if (headerMatch) {
    const countedHashtags = headerMatch[1].length;
    const level = Math.min(countedHashtags, 6);
    tag = `h${level}`;
    content = str.substring(level + 1);
  }
  // Check for list item
  else if (str.startsWith('* ')) {
    tag = 'li';
    content = str.substring(2);
  }
  
  // Process bold emphasis (non-greedy)
  content = content.replace(/\*\*(.+?)\*\*/g, '< strong>$1< /strong>');
  
  return `< ${tag}>${content}< /${tag}>`;
}