export const state = () => ({
  auth: null
})

export const getters = {
  isAuthenticated(state) {
    if (state.auth && state.auth.loggedIn === true) {
      return true
    }
  }
}

export const mutations = {
  setAuth(state, auth) {
    state.auth = auth
  }
}

const cookieparser = process.server ? require('cookieparser') : undefined

export const actions = {
  nuxtServerInit({ commit }, { req }) {
    let auth = null
    if (req.headers.cookie) {
      const parsed = cookieparser.parse(req.headers.cookie)
      try {
        auth = JSON.parse(parsed.auth)
      } catch (err) {
        // No valid cookie found
      }
    }
    commit('setAuth', auth)
  }
}
