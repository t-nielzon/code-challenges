function movingShift(s, shift) {
    // encode the string with increasing shift for each character position
    let encoded = '';
    for (let i = 0; i < s.length; i++) {
        const char = s[i];
        if ((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
            const isLower = char >= 'a';
            const base = isLower ? 'a'.charCodeAt(0) : 'A'.charCodeAt(0);
            const offset = char.charCodeAt(0) - base;
            const newOffset = (offset + shift + i) % 26;
            encoded += String.fromCharCode(base + newOffset);
        } else {
            encoded += char;
        }
    }
    
    // split into 5 parts with optimal distribution
    const length = encoded.length;
    const baseLength = Math.floor(length / 5);
    const numLongerParts = Math.min(4, length - 4 * baseLength);
    
    const parts = [];
    let start = 0;
    
    for (let i = 0; i < 5; i++) {
        let partLength;
        if (i < numLongerParts) {
            partLength = baseLength + 1;
        } else if (i < 4) {
            partLength = baseLength;
        } else {
            partLength = length - start;
        }
        parts.push(encoded.substring(start, start + partLength));
        start += partLength;
    }
    
    return parts;
}

function demovingShift(s, shift) {
    const encoded = s.join('');
    let decoded = '';
    
    for (let i = 0; i < encoded.length; i++) {
        const char = encoded[i];
        if ((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
            const isLower = char >= 'a';
            const base = isLower ? 'a'.charCodeAt(0) : 'A'.charCodeAt(0);
            const offset = char.charCodeAt(0) - base;
            const newOffset = (offset - shift - i + 26 * 100) % 26;
            decoded += String.fromCharCode(base + newOffset);
        } else {
            decoded += char;
        }
    }
    
    return decoded;
}