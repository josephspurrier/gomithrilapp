import m from "mithril";
import Menu from "@/component/menu";
import FlashContainer from "@/component/flashcontainer";

var View = {
  view: function (vnode) {
    return m("main.layout", [
      m(Menu),
      m("section", vnode.children),
      m(FlashContainer),
    ]);
  },
};

export default View;
