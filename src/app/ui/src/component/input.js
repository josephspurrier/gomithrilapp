import m from "mithril"; // eslint-disable-line no-unused-vars

var View = () => {
  return {
    view: (vnode) => (
      <div class="field">
        <label class="label">{vnode.attrs.label}</label>
        <div class="control">
          <input
            name={vnode.attrs.name}
            type={vnode.attrs.type || "text"}
            class="input"
            data-cy={vnode.attrs.name}
            required={vnode.attrs.required}
            oninput={vnode.attrs.oninput}
            value={vnode.attrs.value}
          ></input>
        </div>
      </div>
    ),
  };
};

export default View;
