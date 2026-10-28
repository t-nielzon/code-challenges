function curry(fun, ...args) {
    return function(...newArgs) {
        return fun.apply(this, [...args, ...newArgs]);
    };
}