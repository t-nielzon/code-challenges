function patternNumbers(n) {
    const s = n.toString();
    const patterns = [];
    
    // Match all numbers with fewer digits than n
    for (let len = 1; len < s.length; len++) {
        if (len === 1) {
            patterns.push("[1-9]");
        } else {
            patterns.push("[1-9][0-9]{" + (len - 1) + "}");
        }
    }
    
    // Match numbers with same digit count as n but less than n
    const sameLenPatterns = [];
    let prefix = "";
    
    for (let i = 0; i < s.length; i++) {
        const digit = parseInt(s[i]);
        const remaining = s.length - i - 1;
        
        if (digit > 0) {
            if (i === 0) {
                // First digit must be 1-9, can use digits less than current
                if (digit > 1) {
                    let pattern = "[1-" + (digit - 1) + "]";
                    if (remaining > 0) {
                        pattern += "[0-9]{" + remaining + "}";
                    }
                    sameLenPatterns.push(pattern);
                }
            } else {
                // Middle/last digits can include 0
                let pattern = prefix + "[0-" + (digit - 1) + "]";
                if (remaining > 0) {
                    pattern += "[0-9]{" + remaining + "}";
                }
                sameLenPatterns.push(pattern);
            }
        }
        
        prefix += digit;
    }
    
    if (sameLenPatterns.length > 0) {
        patterns.push(...sameLenPatterns);
    }
    
    // Handle case where no valid numbers exist (n = 1)
    if (patterns.length === 0) {
        return "[^\\s\\S]";
    }
    
    return patterns.join("|");
}