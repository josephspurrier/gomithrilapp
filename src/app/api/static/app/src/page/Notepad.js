import m from "mithril";

var Page = {
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
                />
              </div>
            </div>
          </div>
          <div>
            <ul id="listTodo">
              <li is="note" v-for="(v, k) in todolist"></li>
            </ul>
          </div>
        </div>
      </section>
    ),
};

export default Page;
