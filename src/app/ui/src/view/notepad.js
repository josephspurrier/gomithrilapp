import m from "mithril"; // eslint-disable-line no-unused-vars
import NoteStore from "@/store/notestore";
import Note from "@/component/note";

var Page = {
  oninit: () => {
    NoteStore.load();
  },
  onremove: () => {
    NoteStore.clear();
  },
  view: () => (
    <section id="note-section" class="section">
      <div class="container">
        <div class="box">
          <div class="field">
            <label class="label">To Do</label>
            <div class="control">
              <input
                v-model="inputTodo"
                type="text"
                placeholder="What would you like to do?"
                class="input"
                name="note-add"
                data-cy="note-text"
                onkeypress={(e) => {
                  if (e.key !== "Enter") {
                    return;
                  }
                  NoteStore.submit();
                }}
                oninput={(e) => {
                  NoteStore.current.message = e.target.value;
                }}
                value={NoteStore.current.message}
              />
            </div>
          </div>
          <nav class="level is-mobile">
            <div class="level-left">
              <a title="Add note" class="level-item" onclick={NoteStore.submit}>
                <span class="icon is-small has-text-success">
                  <i class="far fa-plus-square" data-cy="add-note-link"></i>
                </span>
              </a>
            </div>
          </nav>
        </div>
        <div>
          <ul id="listTodo">
            {NoteStore.list.map((note) => (
              <Note
                key={note.id}
                id={note.id}
                message={note.message}
                oninput={(e) => {
                  note.message = e.target.value;
                }}
              />
            ))}
          </ul>
        </div>
      </div>
    </section>
  ),
};

export default Page;
