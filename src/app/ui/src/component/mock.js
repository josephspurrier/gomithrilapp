import m from "mithril";

var Mock = {
  success: (data) => {
    m.mock = {};
    m.request = () =>
      m.mock.request.success
        ? Promise.resolve(m.mock.request)
        : Promise.reject(m.mock.request);
    m.mock.request = data;
    m.mock.request.success = true;
    m.redraw();
  },
};

export default Mock;
