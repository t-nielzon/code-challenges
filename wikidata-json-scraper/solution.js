const https = require('https');

function wikidataScraper(url) {
  return new Promise((resolve, reject) => {
    https.get(url, (res) => {
      let data = '';
      res.on('data', (chunk) => { data += chunk; });
      res.on('end', () => {
        try {
          const json = JSON.parse(data);
          const entityId = Object.keys(json.entities)[0];
          const entity = json.entities[entityId];

          const label = entity.labels && entity.labels.en
            ? entity.labels.en.value
            : 'No Label';

          const description = entity.descriptions && entity.descriptions.en
            ? entity.descriptions.en.value
            : 'No Description';

          resolve({
            ID: entityId,
            LABEL: label,
            DESCRIPTION: description
          });
        } catch (err) {
          reject(err);
        }
      });
    }).on('error', reject);
  });
}