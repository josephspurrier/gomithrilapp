import m from "mithril";
import Menu from "../component/Menu";

var View = {
  view: function (vnode) {
    return m("main.layout", [m(Menu), m("section", vnode.children)]);
  },
};

export default View;
