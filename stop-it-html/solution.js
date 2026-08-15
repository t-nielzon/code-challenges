function loadMyTemplateLanguage() {
  const selfClosingTags = new Set(['link', 'img', 'br', 'hr']);
  const allTags = ['html', 'head', 'title', 'body', 'script', 'link', 'div', 'span', 'img', 'br', 'hr'];
  
  function createElement(tagName, isSelfClosing) {
    return function(attributes = {}, content) {
      let attrString = '';
      for (let [key, value] of Object.entries(attributes)) {
        attrString += ` ${key}="${value}"`;
      }
      
      if (isSelfClosing) {
        return `<${tagName}${attrString} />`;
      } else {
        const contentString = Array.isArray(content) ? content.join('') : '';
        return `<${tagName}${attrString}>${contentString}</${tagName}>`;
      }
    };
  }
  
  allTags.forEach(tag => {
    const isSelfClosing = selfClosingTags.has(tag);
    globalThis[tag] = createElement(tag, isSelfClosing);
  });
}