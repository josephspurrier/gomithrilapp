import m from "mithril";

var User = {
  current: {},
  clear: () => {
    User.current = {};
  },
  register: () => {
    return m.request({
      method: "POST",
      url: "/api/v1/register",
      body: User.current,
    });
  },
  login: () => {
    return m.request({
      method: "POST",
      url: "/api/v1/login",
      body: User.current,
    });
  },
};

export default User;
