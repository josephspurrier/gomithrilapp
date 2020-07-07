import m from "mithril"; // eslint-disable-line no-unused-vars

var Block = () => {
  return {
    view: (vnode) => m("", <div>{vnode.children}</div>),
  };
};

export default Block;
