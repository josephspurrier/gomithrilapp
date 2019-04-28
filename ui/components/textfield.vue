<template>
  <div class="field">
    <label class="label">{{ label }}</label>
    <div class="control">
      <input
        v-validate="{ required }"
        :name="name"
        :type="type"
        :placeholder="placeholder"
        :class="{ input: true, 'is-danger': errors.has(name) }"
        :value="value"
        :disabled="disabled"
        @input="$emit('input', $event.target.value)"
        @keyup.enter="$emit('enter')"
      />
    </div>
    <p v-show="errors.has(name)" class="help is-danger">
      {{ errors.first(name) }}
    </p>
  </div>
</template>

<script>
export default {
  inject: ['$validator'],
  props: {
    label: { type: String, required: true, default: '' },
    name: { type: String, required: true, default: '' },
    placeholder: { type: String, default: '' },
    type: { type: String, default: 'text' },
    value: { type: String, default: '' },
    required: { type: Boolean, default: false },
    disabled: { type: Boolean, default: false }
  }
}
</script>
