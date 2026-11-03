function parityBit(s) {
  return s.split(' ').map(byte => {
    const sevenBits = byte.slice(0, 7);
    const parityBit = parseInt(byte[7]);
    const onesCount = (sevenBits.match(/1/g) || []).length;
    const expectedParity = onesCount % 2;
    
    return parityBit === expectedParity ? sevenBits : 'error';
  }).join(' ');
}