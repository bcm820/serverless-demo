module.exports = {
  qs(event, context) {
    const ints = event.data;
    if (!ints.length) return [];

    const inner = ints => {
      if (!ints.length) return [];
      const [x, ...ys] = ints;
      const lt = inner(ys.filter(y => y <= x));
      const gt = inner(ys.filter(y => y > x));
      return [...lt, x, ...gt];
    };

    return inner(ints);
  }
};
