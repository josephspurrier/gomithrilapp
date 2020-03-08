<template>
  <div class="field">
    <label class="label">{{ label }}</label>
    <div class="control">
      <ValidationProvider v-slot="{ errors }" :rules="isRequired">
        <input
          :name="name"
          :type="type"
          :data-cy="dataCy"
          :placeholder="placeholder"
          :class="{ input: true, 'is-danger': errors.length > 0 }"
          :value="value"
          :disabled="disabled"
          @input="$emit('input', $event.target.value)"
          @keyup.enter="$emit('enter')"
        />
        <p v-show="errors.length > 0" class="help is-danger">
          {{ errors[0] }}
        </p>
      </ValidationProvider>
    </div>
  </div>
</template>

<script>
import { ValidationProvider } from 'vee-validate'

export default {
  components: {
    ValidationProvider
  },
  props: {
    label: { type: String, required: true, default: '' },
    name: { type: String, required: true, default: '' },
    placeholder: { type: String, default: '' },
    type: { type: String, default: 'text' },
    value: { type: String, default: '' },
    required: { type: Boolean, default: false },
    disabled: { type: Boolean, default: false },
    dataCy: { type: String, default: '' }
  },
  computed: {
    isRequired: function() {
      if (this.required) return 'required'
      else return ''
    }
  }
}
</script>
