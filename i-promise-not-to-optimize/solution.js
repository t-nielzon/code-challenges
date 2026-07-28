function antiOptimizeAsync(task) {
    const startTime = Date.now();
    const result = task();
    const elapsed = Date.now() - startTime;
    
    return new Promise(resolve => {
        const targetDelay = 11000 + Math.random() * 1000;
        const remainingDelay = Math.max(0, targetDelay - elapsed);
        
        setTimeout(() => resolve(result), remainingDelay);
    });
}