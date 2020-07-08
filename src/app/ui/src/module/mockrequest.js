import m from "mithril"; // eslint-disable-line no-unused-vars

var Mock = {
  ok: (data, redraw) => {
    Mock.success(data, redraw);
  },
  badRequest: (message, redraw) => {
    Mock.error("Bad Request", message, redraw);
  },
  success: (data, redraw) => {
    m.mock = {};
    m.request = () =>
      m.mock.request.success
        ? Promise.resolve(m.mock.request)
        : Promise.reject(m.mock.request);
    m.mock.request = data;
    m.mock.request.success = true;
    if (redraw) {
      m.redraw();
    }
  },
  error: (status, message, redraw) => {
    m.mock = {};
    m.request = () =>
      m.mock.request.success
        ? Promise.resolve(m.mock.request)
        : Promise.reject(m.mock.request);
    m.mock.request = {
      success: false,
      response: {
        status: status,
        message: message,
      },
    };
    if (redraw) {
      m.redraw();
    }
  },
};

export default Mock;
