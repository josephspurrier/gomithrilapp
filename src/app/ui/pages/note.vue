<template>
  <section class="section">
    <div class="container">
      <div class="box">
        <div class="field">
          <label class="label">To Do</label>
          <div class="control">
            <input
              v-model="inputTodo"
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
            v-for="(v, k) in todolist"
            :id="v.id"
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

export default {
  components: {
    Note
  },
  data() {
    return {
      inputTodo: ''
    }
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
    handleAdd() {
      this.$store.dispatch('todo/addItem', this.inputTodo)
      this.inputTodo = ''
    },
    handleUpdate(index, key, text) {
      this.$store.dispatch('todo/updateItem', { index, key, text })
    },
    handleDelete(index, key) {
      this.$store.dispatch('todo/deleteItem', { index, key })
    }
  }
}
</script>
