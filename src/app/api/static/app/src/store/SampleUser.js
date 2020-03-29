import m from "mithril";

var User = {
  current: {},
  list: [],
  loadList: () => {
    return m
      .request({
        method: "GET",
        url: "https://rem-rest-api.herokuapp.com/api/users",
        withCredentials: true,
      })
      .then((result) => {
        User.list = result.data;
      });
  },
  load: (id) => {
    return m
      .request({
        method: "GET",
        url: "https://rem-rest-api.herokuapp.com/api/users/" + id,
        withCredentials: true,
      })
      .then((result) => {
        User.current = result;
      });
  },
  save: () => {
    return m.request({
      method: "PUT",
      url: "https://rem-rest-api.herokuapp.com/api/users/" + User.current.id,
      withCredentials: true,
      body: User.current,
    });
  },
};

export default User;
