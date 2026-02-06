/*
Extend the `Function` prototype with a method `wrap()` that allows you to wrap an existing function.
The method should behave like this:

function speak(name){
   return "Hello " + name;
}

const wrapper = speak.wrap(function(original, yourName, myName){
   const greeting = original(yourName);
   return greeting + ", my name is " + myName;
});

wrapper("Mary", "Kate");
-->
"Hello Mary, my name is Kate"
*/

Function.prototype.wrap = function(wrapper) {
  // your code here
};