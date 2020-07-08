import m from "mithril"; // eslint-disable-line no-unused-vars
import Flash from "@/component/flash";
import CookieStore from "@/module/cookiestore";

var NoteStore = {
  current: {},
  list: [],
  clear: () => {
    NoteStore.current = {};
  },
  submit: () => {
    NoteStore.create()
      .then(() => {
        Flash.success("Note created.");
        // This could be optimized instead of reloading.
        NoteStore.load();
        NoteStore.clear();
      })
      .catch((err) => {
        Flash.warning(err.response.message);
      });
  },
  create: () => {
    return m.request({
      method: "POST",
      url: "/api/v1/note",
      headers: {
        Authorization: CookieStore.bearerToken(),
      },
      body: NoteStore.current,
    });
  },
  load: () => {
    return m
      .request({
        method: "GET",
        url: "/api/v1/note",
        headers: {
          Authorization: CookieStore.bearerToken(),
        },
      })
      .then((result) => {
        NoteStore.list = result.notes;
      });
  },
  runUpdate: (id, value) => {
    NoteStore.update(id, value).catch((e) => {
      Flash.warning("Could not update note: " + e.response.message);
    });
  },
  update: (id, text) => {
    return m.request({
      method: "PUT",
      url: "/api/v1/note/" + id,
      headers: {
        Authorization: CookieStore.bearerToken(),
      },
      body: { message: text },
    });
  },
  runDelete: (id) => {
    NoteStore.delete(id)
      .then(() => {
        Flash.success("Note deleted.");
        NoteStore.list = NoteStore.list.filter((i) => {
          return i.id !== id;
        });
      })
      .catch((err) => {
        console.log(err);
        Flash.warning("Could not delete: " + err.response.message);
      });
  },
  delete: (id) => {
    return m.request({
      method: "DELETE",
      url: "/api/v1/note/" + id,
      headers: {
        Authorization: CookieStore.bearerToken(),
      },
    });
  },
};

export default NoteStore;
