function countPrimeFactors(n) {
    let count = 0;
    let d = 2;
    while (d * d <= n) {
        while (n % d === 0) {
            n = n / d;
            count++;
        }
        d++;
    }
    if (n > 1) count++;
    return count;
}

function kprimes_step(k, step, start, nd) {
    const kprimes = [];
    for (let i = start; i <= nd; i++) {
        if (countPrimeFactors(i) === k) kprimes.push(i);
    }
    const result = [];
    for (let i = 0; i < kprimes.length - 1; i++) {
        if (kprimes[i + 1] - kprimes[i] === step) {
            result.push([kprimes[i], kprimes[i + 1]]);
        }
    }
    return result;
}