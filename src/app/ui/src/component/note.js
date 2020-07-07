import m from "mithril"; // eslint-disable-line no-unused-vars
import Debounce from "@/module/debounce";
import NoteStore from "@/store/notestore";

var View = {
  oninit: (vnode) => {
    vnode.state.id = vnode.attrs.id;
    vnode.state.message = vnode.attrs.message;
  },
  view: (vnode) => (
    <li style="margin-top: 12px;">
      <div class="box">
        <div class="content">
          <div class="editable">
            <input
              id={vnode.state.id}
              type="text"
              class="input individual-note"
              value={vnode.state.message}
              oninput={(e) => {
                vnode.state.message = e.target.value;
              }}
              onkeyup={(e) => {
                Debounce.run(
                  vnode.state.id,
                  () => {
                    NoteStore.runUpdate(vnode.state.id, e.target.value);
                    vnode.state.saving = "Saving...";
                    m.redraw();
                    setTimeout(function () {
                      vnode.state.saving = "";
                      m.redraw();
                    }, 1000);
                  },
                  1000
                );
              }}
            />
          </div>
        </div>
        <nav class="level is-mobile">
          <div class="level-left">
            <a
              title="Delete note"
              class="level-item"
              onclick={() => {
                NoteStore.runDelete(vnode.state.id);
              }}
            >
              <span class="icon is-small has-text-danger">
                <i class="fas fa-trash" data-cy="delete-note-link"></i>
              </span>
            </a>
          </div>
          <div class="level-right" style="min-height: 1.2rem;">
            <span class="is-size-7 has-text-grey">{vnode.state.saving}</span>
          </div>
        </nav>
      </div>
    </li>
  ),
};

export default View;
