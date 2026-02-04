function computePath(...parts) {
  // Process each part: trim spaces and convert backslashes to slashes
  const processedParts = parts
    .map(part => part.trim().replace(/\\/g, '/'))
    .filter(part => part !== '');
  
  // Join all parts with slash
  let path = processedParts.join('/');
  
  // Replace multiple consecutive slashes with single slash
  path = path.replace(/\/+/g, '/');
  
  // Remove trailing slash (but keep if it's the only character)
  if (path.length > 1 && path.endsWith('/')) {
    path = path.slice(0, -1);
  }
  
  // Empty path becomes a slash
  if (path === '') {
    path = '/';
  }
  
  return path;
}