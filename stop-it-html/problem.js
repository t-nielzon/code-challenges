// "Stop it, HTML!"
// 
// You don't quite like HTML. In fact, you like what HTML does but you find annoying to write HTML, if not even expensive!
// So you decided to explore new ways to write YOUR template language.
//
// After some times you came up with a JS template you think it could be a good start.
//
// element(attributes, content)
//
// Where `element` returns a `string` and has the following parameters:
// - attributes: object (required, can be empty) - contains key: value pairs
// - content: array (optional, can be undefined) - contains strings and other element results
//
// You need to support the following tags:
// ['html','head','title','body','script','link','div','span','img','br','hr']
//
// Some tags (link, img, br, hr) are self-closing and should use the /> format.
// 
// loadMyTemplateLanguage() is a setup function that will be called before any assertions.
// You may use it to attach these functions to the global scope.

function loadMyTemplateLanguage() {
  
}