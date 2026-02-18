var Format = (function() {
  const tags = ['div', 'p', 'span', 'h1'];

  function createFormatter(chain) {
    // callable function that wraps args in the chained tags
    const fn = function(...args) {
      const content = args.join('');
      return chain.reduceRight((inner, tag) => `<${tag}>${inner}</${tag}>`, content);
    };

    // each tag property returns a new formatter with that tag appended to the chain
    tags.forEach(tag => {
      Object.defineProperty(fn, tag, {
        get() {
          return createFormatter([...chain, tag]);
        }
      });
    });

    return fn;
  }

  const format = {};
  tags.forEach(tag => {
    Object.defineProperty(format, tag, {
      get() {
        return createFormatter([tag]);
      }
    });
  });

  return format;
})();