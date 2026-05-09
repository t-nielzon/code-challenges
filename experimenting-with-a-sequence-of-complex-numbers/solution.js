function f(x, y, eps) {
    const r = Math.hypot(x, y);
    if (r >= 1) return -1;
    const v = Math.log(eps) / Math.log(r);
    let np1 = Math.floor(v) + 1;
    if (np1 <= v) np1 += 1;
    return np1 - 1;
}