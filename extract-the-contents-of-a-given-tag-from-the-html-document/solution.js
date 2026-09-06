function getTagContent(html, tag) {
  const escapedTag = tag.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  const regex = new RegExp(`<${escapedTag}(?:\\s[^>]*)?>([\\s\\S]*?)</${escapedTag}>`, 'g');
  const matches = [];
  let match;
  
  while ((match = regex.exec(html)) !== null) {
    matches.push(match[1].trim());
  }
  
  return matches;
}