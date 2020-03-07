<template>
  <div>
    <section class="section">
      <div class="container">
        <h1 class="title">{{ title }}</h1>
        <h2 class="subtitle">{{ subtitle }}</h2>
      </div>

      <div class="container" style="margin-top: 1em;">
        <form name="register">
          <textfield
            v-model="register.first_name"
            :disabled="isLoading"
            label="First Name"
            name="first_name"
            type="text"
            data-cy="first_name"
            required
          ></textfield>
          <textfield
            v-model="register.last_name"
            :disabled="isLoading"
            label="Last Name"
            name="last_name"
            type="text"
            data-cy="last_name"
            required
          ></textfield>
          <textfield
            v-model="register.email"
            :disabled="isLoading"
            label="Email"
            name="email"
            type="text"
            data-cy="email"
            required
          ></textfield>
          <textfield
            v-model="register.password"
            :disabled="isLoading"
            @enter="submit"
            label="Password"
            name="password"
            type="password"
            data-cy="password"
            required
          ></textfield>
          <div class="field is-grouped">
            <p class="control">
              <a
                id="submit"
                :class="{
                  button: true,
                  'is-primary': true,
                  'is-loading': isLoading
                }"
                @click="submit"
                data-cy="submit"
                >Create Account</a
              >
            </p>
            <p class="control">
              <a
                :class="{
                  button: true,
                  'is-light': true,
                  'is-loading': isLoading
                }"
                @click="clear"
                >Clear</a
              >
            </p>
          </div>
        </form>
      </div>
    </section>
  </div>
</template>

<script>
import textfield from '~/components/textfield.vue'
import { HTTP } from '~/modules/http-common'
import Flash from '~/modules/flash.js'

export default {
  components: { textfield },
  $validates: true,
  data() {
    return {
      title: 'Register',
      subtitle: 'Enter your information below.',
      register: {
        first_name: '',
        last_name: '',
        email: '',
        password: ''
      },
      isLoading: false
    }
  },
  head() {
    return {
      title: this.title
    }
  },
  middleware: 'notAllowIfAuthenticated',
  methods: {
    // clear will clear the form.
    clear() {
      this.register.first_name = ''
      this.register.last_name = ''
      this.register.email = ''
      this.register.password = ''
    },
    submit() {
      // Create the flash object.
      const f = new Flash()

      // Validate the form.
      this.$validator
        .validateAll()
        .then(result => {
          if (result === true) {
            this.submitReady()
          }
        })
        .catch(() => {
          f.warning('Could not validate the form.')
        })
    },
    // submit will send a register request to the server.
    submitReady() {
      // Create the flash object.
      const f = new Flash()
      this.isLoading = true
      let success = false

      // Send a register request to the server.
      const headers = {
        headers: {
          'Content-Type': 'application/json'
        }
      }
      HTTP.post(`v1/register`, this.register, headers)
        .then(response => {
          if (response.data.status === 'Created') {
            f.success('Registered.')
            success = true
          } else {
            f.failed('Response received is not in the correct format.')
          }
        })
        .catch(err => {
          if (err.response === undefined) {
            f.warning(
              'There was an error reaching the server. Please try again later.' +
                err
            )
          } else if (err.response.data.message !== undefined) {
            f.warning(err.response.data.message)
          } else {
            f.warning('There was an error. Please try again later.' + err)
          }
        })
        .finally(() => {
          this.clear()
          this.isLoading = false

          if (success) {
            this.$router.push('/')
          }
        })
    }
  }
}
</script>
