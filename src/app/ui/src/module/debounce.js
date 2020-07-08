var Debounce = {
  m: new Map(),
  run(id, func, timeout) {
    if (Debounce.m[id]) {
      clearTimeout(Debounce.m[id]);
    }
    Debounce.m[id] = setTimeout(func, timeout);
  },
};

export default Debounce;
