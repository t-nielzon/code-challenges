// Limit Async Tasks (5 kyu)
//
// Running asynchronous work with a concurrency limit is a common scheduling problem.
// Sometimes you need to execute many tasks, but you cannot allow all of them to run at once.
//
// Build a function that runs asynchronous tasks while respecting a maximum number of tasks
// that may run at the same time.
//
// Given:
// - tasks: an array of functions
// - limit: the maximum number of tasks that may run simultaneously
//
// When there is available capacity and there are still pending tasks, another task must be
// started immediately.
//
// Each function in tasks is called only when it is scheduled to run.
// A task normally produces its result asynchronously, but it may also fail immediately by
// throwing an error.
//
// Return an asynchronous result that:
// - succeeds with an array of results in the same order as the input tasks
// - fails with the same error if any task fails
//
// If limit is less than 1, fail with an error whose message is "Invalid limit"

function limitAsyncTasks(tasks, limit) {
  
}