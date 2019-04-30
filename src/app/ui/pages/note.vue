<template>
  <section class="section">
    <div class="container">
      <div class="box">
        <div class="field">
          <label class="label">To Do</label>
          <div class="control">
            <input
              id="inputTodo"
              class="input"
              type="text"
              placeholder="What would you like to do?"
              @keydown.enter="handleAdd"
            />
          </div>
        </div>
      </div>
      <div>
        <ul id="listTodo">
          <li
            is="Note"
            v-for="(v, k) in TodoList.list"
            :key="v.id"
            :index="k"
            :message="v.message"
            @remove="handleDelete"
            @edit="handleUpdate"
          ></li>
        </ul>
      </div>
    </div>
  </section>
</template>

<script>
import Note from '~/components/Note.vue'
import TDD from '~/modules/todo.js'

export default {
  components: {
    Note
  },
  data() {
    return {
      TodoList: TDD
    }
  },
  mounted() {
    this.TodoList.loadItems()
    // console.log('mounted')
  },
  methods: {
    handleAdd() {
      this.TodoList.addItem()
    },
    handleUpdate(id, text) {
      this.TodoList.updateItem(id, text)
    },
    handleDelete(id) {
      this.TodoList.deleteItem(id)
    }
  }
}
</script>
