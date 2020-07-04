// eslint-disable-next-line no-unused-vars
import m from "mithril";
import Flash from "./flash";

var View = {
  view: () => (
    <div
      id={Flash.identifier}
      style="position: fixed; bottom: 1.5rem; right: 1.5rem; z-index: 100; margin: 0;"
    ></div>
  ),
};

export default View;
