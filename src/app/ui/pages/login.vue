<template>
  <div>
    <section class="section">
      <div class="container">
        <h1 class="title">
          Login
        </h1>
        <h2 class="subtitle">
          Enter your login information below.
        </h2>
      </div>

      <div class="container" style="margin-top: 1em;">
        <form name="login">
          <textfield
            v-model="login.email"
            label="Email"
            name="email"
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
                >Submit</a
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
            <p class="control">
              <a
                :class="{
                  button: true,
                  'is-light': true,
                  'is-loading': isLoading
                }"
                :href="register"
                >Register</a
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
        email: '',
        password: ''
      },
      isLoading: false,
      register: '/register'
    }
  },
  middleware: 'notAllowIfAuthenticated',
  methods: {
    // clear will clear the form.
    clear() {
      this.login.email = ''
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
      HTTP.post(`login`, this.login)
        .then(response => {
          if (response.data.token !== undefined) {
            f.success('Login successful.')
            const auth = {
              accessToken: response.data.token,
              loggedIn: true
            }
            this.$store.commit('setAuth', auth) // mutating to store for client rendering
            Cookie.set('auth', auth) // saving token in cookie for server rendering
            success = true
          } else {
            f.failed('Token is not in the correct format.')
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
