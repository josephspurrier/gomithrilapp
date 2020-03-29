import m from "mithril";
import Cookie from "js-cookie";

function getBearer() {
  let auth = Cookie.get("auth");
  if (auth === undefined) {
    return false;
  }

  let v = JSON.parse(auth);
  return "Bearer " + v.accessToken;
}

var Note = {
  current: {},
  list: [],
  clear: () => {
    Note.current = {};
  },
  create: () => {
    return m.request({
      method: "POST",
      url: "/api/v1/note",
      headers: {
        Authorization: getBearer(),
      },
      body: Note.current,
    });
  },
  load: () => {
    return m
      .request({
        method: "GET",
        url: "/api/v1/note",
        headers: {
          Authorization: getBearer(),
        },
      })
      .then((result) => {
        Note.list = result.notes;
        console.log(Note.list);
      });
  },
  update: (id, text) => {
    return m.request({
      method: "PUT",
      url: "/api/v1/note/" + id,
      headers: {
        Authorization: getBearer(),
      },
      body: { message: text },
    });
  },
  delete: (id) => {
    return m.request({
      method: "DELETE",
      url: "/api/v1/note/" + id,
      headers: {
        Authorization: getBearer(),
      },
    });
  },
};

export default Note;
