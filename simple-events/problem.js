/*
 * Your goal is to write an Event constructor function, which can be used to make event objects.
 * 
 * An event object should work like this:
 *   - it has a .subscribe() method, which takes a function and stores it as its handler
 *   - it has an .unsubscribe() method, which takes a function and removes it from its handlers
 *   - it has an .emit() method, which takes an arbitrary number of arguments and calls all the stored functions with these arguments
 * 
 * Simplifications:
 *   - all functions are called with correct arguments (only functions will be passed to unsubscribe)
 *   - you should not worry about the order of handlers' execution
 *   - the handlers will not attempt to modify an event object
 *   - the context of handlers' execution is not important
 *   - each handler will be subscribed at most once at any given moment of time
 */

function Event() {
  
}