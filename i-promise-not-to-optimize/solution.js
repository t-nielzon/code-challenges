function antiOptimizeAsync(task) {
    // capture the start time so total elapsed can be measured regardless
    // of how long the (0-10s) task itself takes to run.
    const start = Date.now();

    // run the task on its own so its variable runtime doesn't eat into
    // the guaranteed delay window; the result is resolved only after the
    // 11-12 second window has fully elapsed.
    const result = task();

    return new Promise(resolve => {
        const remaining = 11000 - (Date.now() - start);
        setTimeout(() => resolve(result), Math.max(0, remaining));
    });
}