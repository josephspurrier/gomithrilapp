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
            >
          </div>
        </div>
      </div>
      <div>
        <ul id="listTodo">
          <li
            is="Note"
            v-for="(v, k) in todolist"
            :key="v.id"
            :index="k"
            :message="v.message"
            :id="v.id"
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

export default {
  components: {
    Note
  },
  computed: {
    todolist() {
      return this.$store.state.todo.list
    }
  },
  mounted() {
    this.$store.dispatch('todo/loadItems')
  },
  methods: {
    handleAdd(id) {
      this.$store.dispatch('todo/addItem', inputTodo.value)
      inputTodo.value = ''
    },
    handleUpdate(index, id, text) {
      this.$store.dispatch('todo/updateItem', { index, id, text })
    },
    handleDelete(index, id) {
      this.$store.dispatch('todo/deleteItem', { index, id })
    }
  }
}
</script>
