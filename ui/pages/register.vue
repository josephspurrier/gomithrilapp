<template>
  <div>
    <section class="section">
      <div class="container">
        <h1 class="title">
          Register
        </h1>
        <h2 class="subtitle">
          Enter your information below.
        </h2>
      </div>

      <div class="container" style="margin-top: 1em;">
        <form name="login">
          <textfield
            v-model="login.username"
            label="Username"
            name="username"
            type="text"
            :disabled="isLoading"
            required
            @enter="submit"
          ></textfield>
          <textfield
            v-model="login.password"
            label="Password"
            name="password"
            type="password"
            :disabled="isLoading"
            required
            @enter="submit"
          ></textfield>
          <div class="field is-grouped">
            <p class="control">
              <a
                :class="{
                  button: true,
                  'is-primary': true,
                  'is-loading': isLoading
                }"
                @click="submit"
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

const Cookie = process.client ? require('js-cookie') : undefined

export default {
  components: { textfield },
  $validates: true,
  data() {
    return {
      login: {
        username: '',
        password: ''
      },
      isLoading: false
    }
  },
  middleware: 'notAllowIfAuthenticated',
  methods: {
    // clear will clear the form.
    clear() {
      this.login.username = ''
      this.login.password = ''
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
    // submit will send a login request to the server.
    submitReady() {
      // Create the flash object.
      const f = new Flash()
      this.isLoading = true
      let success = false

      // Send a login request to the server.
      HTTP.post(`register`, this.login)
        .then(response => {
          if (response.data.success === true) {
            f.success('Registered.')
            success = true
          } else if (response.data.success === undefined) {
            f.failed('Response received is not in the correct format.')
          } else {
            f.failed('User already exists.')
          }
        })
        .catch(err => {
          f.warning('There was an error. Please try again later.' + err)
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
