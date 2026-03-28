function diag2Sym(s) {
    const rows = s.split('\n');
    const n = rows.length;
    const result = [];
    for (let i = n - 1; i >= 0; i--) {
        let line = '';
        for (let j = n - 1; j >= 0; j--) {
            line += rows[j][i];
        }
        result.push(line);
    }
    return result.join('\n');
}

function rot90Counter(s) {
    const rows = s.split('\n');
    const n = rows.length;
    const result = [];
    for (let i = n - 1; i >= 0; i--) {
        let line = '';
        for (let j = 0; j < n; j++) {
            line += rows[j][i];
        }
        result.push(line);
    }
    return result.join('\n');
}

function selfieDiag2Counterclock(s) {
    const orig = s.split('\n');
    const d = diag2Sym(s).split('\n');
    const r = rot90Counter(s).split('\n');
    return orig.map((row, i) => row + '|' + d[i] + '|' + r[i]).join('\n');
}

function oper(fct, s) {
    return fct(s);
}