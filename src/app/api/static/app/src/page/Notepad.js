import m from "mithril";
import Note from "../store/Note";
import Flash from "../module/Flash";

function onSubmit() {
  Note.create()
    .then(() => {
      Flash.success("Note created.");
      // This could be optimized instead of reloading.
      Note.load();
      Note.clear();
    })
    .catch((err) => {
      Flash.warning(err.response.message);
    });
}

var Page = {
  oninit: () => {
    Note.load();
  },
  view: () =>
    m(
      "notepad",
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
                    onSubmit();
                  }}
                  oninput={(e) => {
                    Note.current.message = e.target.value;
                  }}
                  value={Note.current.message}
                />
              </div>
            </div>
            <nav class="level is-mobile">
              <div class="level-left">
                <a title="Add note" class="level-item" onclick={onSubmit}>
                  <span class="icon is-small has-text-success">
                    <i class="far fa-plus-square" data-cy="add-note-link"></i>
                  </span>
                </a>
              </div>
            </nav>
          </div>
          <div>
            <ul id="listTodo">
              {Note.list.map((note) => (
                <li key={note.id}>
                  <div class="box">
                    <div class="content">
                      <div class="editable">
                        <input
                          id={note.id}
                          type="text"
                          class="input individual-note"
                          value={note.message}
                          oninput={(e) => {
                            note.message = e.target.value;
                          }}
                          onkeyup={(e) => {
                            //note.message = e.target.value;
                            console.log("keypress", e.target.value);
                            Note.update(note.id, e.target.value)
                              .then(() => {
                                Flash.success("Note updated.");
                              })
                              .catch((e) => {
                                Flash.warning(
                                  "Could not update note: " + e.response.message
                                );
                              });
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
                            Note.delete(note.id)
                              .then(() => {
                                Flash.success("Note deleted.");
                                Note.list = Note.list.filter(function (i) {
                                  return i.id !== note.id;
                                });
                              })
                              .catch((err) => {
                                console.log(err);
                                Flash.warning(
                                  "Could not delete: " + err.response.message
                                );
                              });
                          }}
                        >
                          <span class="icon is-small has-text-danger">
                            <i
                              class="fas fa-trash"
                              data-cy="delete-note-link"
                            ></i>
                          </span>
                        </a>
                      </div>
                    </nav>
                  </div>
                </li>
              ))}
            </ul>
          </div>
        </div>
      </section>
    ),
};

export default Page;
