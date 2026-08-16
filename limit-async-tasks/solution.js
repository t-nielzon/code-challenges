function limitAsyncTasks(tasks, limit) {
  if (limit < 1) {
    return Promise.reject(new Error('Invalid limit'));
  }
  
  if (tasks.length === 0) {
    return Promise.resolve([]);
  }
  
  return new Promise((resolve, reject) => {
    const results = new Array(tasks.length);
    let running = 0;
    let completed = 0;
    let nextIndex = 0;
    let rejected = false;
    
    function processTask(index) {
      running++;
      
      try {
        Promise.resolve(tasks[index]())
          .then(result => {
            if (rejected) return;
            results[index] = result;
            completed++;
            running--;
            
            if (completed === tasks.length) {
              resolve(results);
            } else {
              scheduleNext();
            }
          })
          .catch(error => {
            if (!rejected) {
              rejected = true;
              reject(error);
            }
          });
      } catch (error) {
        if (!rejected) {
          rejected = true;
          reject(error);
        }
      }
    }
    
    function scheduleNext() {
      if (rejected) return;
      while (running < limit && nextIndex < tasks.length) {
        processTask(nextIndex++);
      }
    }
    
    scheduleNext();
  });
}