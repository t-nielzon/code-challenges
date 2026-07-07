async function scrapeWikidataJson(url) {
  const response = await fetch(url);
  const data = await response.json();
  
  // Extract ID from URL
  const idMatch = url.match(/Q\d+/);
  const id = idMatch ? idMatch[0] : "Unknown";
  
  // Get the entity from data
  const entity = data.entities && data.entities[id];
  
  // Initialize defaults
  let label = "No Label";
  let description = "No Description";
  
  if (entity) {
    // Get English label (exact "en" key only)
    if (entity.labels && entity.labels.en && entity.labels.en.value) {
      label = entity.labels.en.value;
    }
    
    // Get English description (exact "en" key only)
    if (entity.descriptions && entity.descriptions.en && entity.descriptions.en.value) {
      description = entity.descriptions.en.value;
    }
  }
  
  return {
    "ID": id,
    "LABEL": label,
    "DESCRIPTION": description
  };
}