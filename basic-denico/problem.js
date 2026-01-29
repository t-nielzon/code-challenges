/*
# Task

Write a function `deNico` that accepts two parameters:
- `key` - string consists of unique letters and digits
- `message` - string with encoded message 

and decodes the `message` using the `key`.  

First create a numeric key basing on the provided `key` by assigning each letter position in which it is located after setting the letters from `key` in an alphabetical order.

For example, for the key `crazy` we will get `23154` because of `acryz` (sorted letters from the key).  
Let's decode  `cseerntiofarmit on  ` using our `crazy` key.
