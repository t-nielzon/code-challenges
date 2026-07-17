function loadMyTemplateLanguage() {
  const selfClosingTags = new Set(['link', 'img', 'br', 'hr']);
  const tags = ['html', 'head', 'title', 'body', 'script', 'link', 'div', 'span', 'img', 'br', 'hr'];
  
  function createTag(tagName) {
    return function(attributes = {}, content) {
      let html = `<${tagName}`;
      
      for (const [key, value] of Object.entries(attributes)) {
        html += ` ${key}="${value}"`;
      }
      
      if (selfClosingTags.has(tagName)) {
        html += ' />';
      } else {
        html += '>';
        if (content) {
          for (const item of content) {
            html += item;
          }
        }
        html += `</${tagName}>`;
      }
      
      return html;
    };
  }
  
  for (const tag of tags) {
    global[tag] = createTag(tag);
  }
}