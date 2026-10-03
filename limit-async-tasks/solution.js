function limitAsyncTasks(tasks, limit) {
  if (limit < 1) {
    return Promise.reject(new Error("Invalid limit"));
  }
  
  if (tasks.length === 0) {
    return Promise.resolve([]);
  }
  
  return new Promise((resolve, reject) => {
    const results = new Array(tasks.length);
    let completed = 0;
    let running = 0;
    let currentIndex = 0;
    let failed = false;
    
    function tryStartTasks() {
      if (failed) return;
      
      while (running < limit && currentIndex < tasks.length) {
        const index = currentIndex++;
        running++;
        
        try {
          const result = tasks[index]();
          
          Promise.resolve(result).then(
            (value) => {
              results[index] = value;
              completed++;
              running--;
              
              if (completed === tasks.length) {
                resolve(results);
              } else {
                tryStartTasks();
              }
            },
            (error) => {
              if (!failed) {
                failed = true;
                reject(error);
              }
            }
          );
        } catch (error) {
          if (!failed) {
            failed = true;
            reject(error);
          }
          return;
        }
      }
    }
    
    tryStartTasks();
  });
}