import m from "mithril";
import Auth from "../module/auth";

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
        Authorization: Auth.bearerToken(),
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
          Authorization: Auth.bearerToken(),
        },
      })
      .then((result) => {
        Note.list = result.notes;
      });
  },
  update: (id, text) => {
    return m.request({
      method: "PUT",
      url: "/api/v1/note/" + id,
      headers: {
        Authorization: Auth.bearerToken(),
      },
      body: { message: text },
    });
  },
  delete: (id) => {
    return m.request({
      method: "DELETE",
      url: "/api/v1/note/" + id,
      headers: {
        Authorization: Auth.bearerToken(),
      },
    });
  },
};

export default Note;
