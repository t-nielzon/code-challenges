function minCoffeeSwaps(tray, deliverList) {
    function rotate(m) {
        const r = m.length, c = m[0].length;
        const result = [];
        for (let i = 0; i < c; i++) {
            const row = [];
            for (let j = 0; j < r; j++) row.push(m[r - 1 - j][i]);
            result.push(row);
        }
        return result;
    }

    function mirror(m) {
        return m.map(row => [...row].reverse());
    }

    function flatten(m) {
        const out = [];
        for (const row of m) for (const x of row) out.push(x);
        return out;
    }

    function minSwaps(arr, target) {
        const pos = new Map();
        target.forEach((v, i) => pos.set(v, i));
        const n = arr.length;
        const visited = new Array(n).fill(false);
        let cycles = 0;
        for (let i = 0; i < n; i++) {
            if (visited[i]) continue;
            cycles++;
            let j = i;
            while (!visited[j]) {
                visited[j] = true;
                j = pos.get(arr[j]);
            }
        }
        return n - cycles;
    }

    const orientations = [];
    let cur = tray;
    for (let i = 0; i < 4; i++) {
        orientations.push(cur);
        orientations.push(mirror(cur));
        cur = rotate(cur);
    }

    let best = Infinity;
    for (const o of orientations) {
        const flat = flatten(o);
        const swaps = minSwaps(flat, deliverList);
        if (swaps < best) best = swaps;
    }
    return best;
}