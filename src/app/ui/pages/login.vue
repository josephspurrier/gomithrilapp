<template>
  <div>
    <section class="section">
      <div class="container">
        <h1 class="title">{{ title }}</h1>
        <h2 class="subtitle">{{ subtitle }}</h2>
      </div>

      <div class="container" style="margin-top: 1em;">
        <ValidationObserver ref="observer">
          <form name="login" @submit.prevent="onSubmit">
            <textfield
              v-model="login.email"
              :disabled="isLoading"
              label="Email"
              name="email"
              type="text"
              data-cy="email"
              required
            ></textfield>
            <textfield
              v-model="login.password"
              :disabled="isLoading"
              label="Password"
              name="password"
              type="password"
              data-cy="password"
              required
            ></textfield>
            <div class="field is-grouped">
              <p class="control">
                <button
                  id="submit"
                  :class="{
                    button: true,
                    'is-primary': true,
                    'is-loading': isLoading
                  }"
                  type="submit"
                  data-cy="submit"
                >
                  Submit
                </button>
              </p>
              <p class="control">
                <button
                  :class="{
                    button: true,
                    'is-light': true,
                    'is-loading': isLoading
                  }"
                  type="button"
                  @click="clear"
                >
                  Clear
                </button>
              </p>
              <p class="control">
                <n-link
                  :class="{
                    button: true,
                    'is-light': true,
                    'is-loading': isLoading
                  }"
                  to="register"
                >
                  Register
                </n-link>
              </p>
            </div>
          </form>
        </ValidationObserver>
      </div>
    </section>
  </div>
</template>

<script>
import { ValidationObserver } from 'vee-validate'
import textfield from '~/components/textfield.vue'
import { HTTP } from '~/modules/http-common'
import Flash from '~/modules/flash.js'

const Cookie = process.client ? require('js-cookie') : undefined

export default {
  components: {
    textfield,
    ValidationObserver
  },
  data() {
    return {
      title: 'Login',
      subtitle: 'Enter your login information below.',
      count: 0,
      login: {
        email: '',
        password: ''
      },
      isLoading: false,
      register: '/register',
      lastError: ''
    }
  },
  methods: {
    // clear will clear the form.
    clear() {
      this.login.email = ''
      this.login.password = ''

      // Reset the form so there are no errors.
      this.$refs.observer.reset()
    },
    async onSubmit() {
      // Create the flash object.
      const f = new Flash()

      const isValid = await this.$refs.observer.validate()
      if (isValid) {
        this.submitReady()
      } else {
        f.warning('One or more required fields is missing.')
      }
    },
    // submit will send a login request to the server.
    submitReady() {
      // Create the flash object.
      const f = new Flash()
      this.isLoading = true
      let success = false
      this.count += 1

      // Send a login request to the server.
      const headers = {
        headers: {
          'Content-Type': 'application/json'
        }
      }
      HTTP.post(`v1/login`, this.login, headers)
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
  },
  head() {
    return {
      title: this.title
    }
  },
  middleware: 'notAllowIfAuthenticated'
}
</script>
