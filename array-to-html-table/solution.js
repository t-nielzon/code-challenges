function toTable(data, headers = false, index = false) {
  let html = '<table>';
  
  let startRow = 0;
  
  if (headers) {
    html += '<thead><tr>';
    
    if (index) {
      html += '<th></th>';
    }
    
    for (let cell of data[0]) {
      html += '<th>' + (cell === null ? '' : String(cell)) + '</th>';
    }
    
    html += '</tr></thead>';
    startRow = 1;
  }
  
  html += '<tbody>';
  
  for (let i = startRow; i < data.length; i++) {
    html += '<tr>';
    
    if (index) {
      html += '<td>' + (i - startRow + 1) + '</td>';
    }
    
    for (let cell of data[i]) {
      html += '<td>' + (cell === null ? '' : String(cell)) + '</td>';
    }
    
    html += '</tr>';
  }
  
  html += '</tbody></table>';
  
  return html;
}