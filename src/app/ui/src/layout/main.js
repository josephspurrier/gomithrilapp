import m from "mithril"; // eslint-disable-line no-unused-vars
import Menu from "@/component/menu";
import Flash from "@/component/flash";

var View = {
  view: function (vnode) {
    return m("main.layout", [m(Menu), m("section", vnode.children), m(Flash)]);
  },
};

export default View;
