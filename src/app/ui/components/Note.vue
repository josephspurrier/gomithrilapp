<template>
  <li>
    <div class="box">
      <div class="content">
        <div class="editable">
          <input
            :id="id"
            v-model="imessage"
            type="text"
            class="input individual-note"
            @keyup="$emit('edit', index, id, imessage)"
          />
        </div>
      </div>
      <nav class="level is-mobile">
        <div class="level-left">
          <a class="level-item">
            <span class="icon is-small has-text-danger">
              <i
                class="fa fa-trash-o"
                data-cy="delete-note-link"
                @click="$emit('remove', index, id)"
              ></i>
            </span>
          </a>
        </div>
      </nav>
    </div>
  </li>
</template>

<script>
export default {
  props: {
    index: {
      type: Number,
      required: true
    },
    id: {
      type: String,
      required: true
    },
    message: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      imessage: this.message
    }
  }
}
</script>
