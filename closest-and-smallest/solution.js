function closest(strng) {
    if (!strng || strng.trim() === "") return [];
    
    const numbers = strng.split(" ").map(Number);
    
    // calculate weight of each number
    const weights = numbers.map(num => {
        return String(num).split("").reduce((sum, digit) => sum + Number(digit), 0);
    });
    
    // find minimum difference in weights
    let minDiff = Infinity;
    for (let i = 0; i < numbers.length; i++) {
        for (let j = i + 1; j < numbers.length; j++) {
            minDiff = Math.min(minDiff, Math.abs(weights[i] - weights[j]));
        }
    }
    
    // collect all pairs with minimum difference
    const candidates = [];
    for (let i = 0; i < numbers.length; i++) {
        for (let j = i + 1; j < numbers.length; j++) {
            if (Math.abs(weights[i] - weights[j]) === minDiff) {
                candidates.push([i, j, weights[i], weights[j], numbers[i], numbers[j]]);
            }
        }
    }
    
    // sort by: smallest weight, then largest weight, then smallest index, then largest index
    candidates.sort((a, b) => {
        const [i_a, j_a, w_i_a, w_j_a] = a;
        const [i_b, j_b, w_i_b, w_j_b] = b;
        
        const min_w_a = Math.min(w_i_a, w_j_a);
        const max_w_a = Math.max(w_i_a, w_j_a);
        const min_w_b = Math.min(w_i_b, w_j_b);
        const max_w_b = Math.max(w_i_b, w_j_b);
        
        if (min_w_a !== min_w_b) return min_w_a - min_w_b;
        if (max_w_a !== max_w_b) return max_w_a - max_w_b;
        
        const min_i_a = Math.min(i_a, j_a);
        const max_i_a = Math.max(i_a, j_a);
        const min_i_b = Math.min(i_b, j_b);
        const max_i_b = Math.max(i_b, j_b);
        
        if (min_i_a !== min_i_b) return min_i_a - min_i_b;
        return max_i_a - max_i_b;
    });
    
    const [i, j, w_i, w_j, n_i, n_j] = candidates[0];
    
    // format output: sort by weight, then by index if weights are equal
    if (w_i < w_j || (w_i === w_j && i < j)) {
        return [[w_i, i, n_i], [w_j, j, n_j]];
    } else {
        return [[w_j, j, n_j], [w_i, i, n_i]];
    }
}