import m from "mithril";

var User = {
  register: function () {
    return m.request({
      method: "POST",
      url: "/test",
      withCredentials: true,
    });
  },
};

export default User;
