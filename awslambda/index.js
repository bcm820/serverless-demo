exports.handler = function(event, context) {
  const { unsorted } = event;
  if (!unsorted || !unsorted.length) return context.done(null, []);

  const quickSort = unsorted => {
    if (!unsorted.length) return [];
    const [x, ...ys] = unsorted;
    const lt = quickSort(ys.filter(y => y <= x));
    const gt = quickSort(ys.filter(y => y > x));
    return [...lt, x, ...gt];
  };

  return context.done(null, quickSort(unsorted));
};
