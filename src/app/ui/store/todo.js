import shortid from 'shortid'
import { HTTP } from '~/modules/http-common'
import Flash from '~/modules/flash.js'

const STORAGE_KEY = 'app-note'

export const state = () => ({
  list: [],
})

export const actions = {
  addItem: function (state, text) {
    let accessToken = ''
    if (state.rootState.auth && state.rootState.auth.loggedIn) {
      accessToken = state.rootState.auth.accessToken
    }

    const f = new Flash()
    let success = false

    // Send a login request to the server.
    const headers = {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + accessToken,
      }
    }

    let form = {
      message: text,
    }

    HTTP.post(`v1/note`, form, headers)
      .then(response => {
        if (response.data !== undefined) {
          f.success('Created.')
          state.commit('addListItem', {
            id: response.data.record_id || shortid.generate(),
            message: text,
          })
          state.commit('save')
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
        if (!success) {
        }
      })
  },
  updateItem: function (state, { index, id, text }) {
    let accessToken = ''
    if (state.rootState.auth && state.rootState.auth.loggedIn) {
      accessToken = state.rootState.auth.accessToken
    }

    const f = new Flash()
    let success = false

    // Send a login request to the server.
    const headers = {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + accessToken,
      }
    }

    let form = {
      message: text,
    }

    HTTP.put(`v1/note/` + id, form, headers)
      .then(response => {
        if (response.data !== undefined) {
          f.success('Saved.')
          state.commit('updateItem', { index, text })
          state.commit('save')
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
        if (!success) {
        }
      })
  },
  deleteItem: function (state, { index, id }) {
    let accessToken = ''
    if (state.rootState.auth && state.rootState.auth.loggedIn) {
      accessToken = state.rootState.auth.accessToken
    }

    const f = new Flash()
    let success = false

    // Send a login request to the server.
    const headers = {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + accessToken,
      }
    }
    HTTP.delete(`v1/note/` + id, headers)
      .then(response => {
        if (response.data !== undefined) {
          f.success('Deleted.')
          state.commit('deleteItem', index)
          state.commit('save')
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
        if (!success) {
        }
      })
  },
  loadItems: function (state) {
    let accessToken = ''
    if (state.rootState.auth && state.rootState.auth.loggedIn) {
      accessToken = state.rootState.auth.accessToken
    }

    const f = new Flash()
    let success = false

    // Send a login request to the server.
    const headers = {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + accessToken,
      }
    }
    HTTP.get(`v1/note`, headers)
      .then(response => {
        if (response.data !== undefined) {
          f.success('Data loaded.')
          state.commit('setItems', response.data.notes)
          state.commit('save')
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
        if (!success) {
        }
      })
  },
}

export const mutations = {
  save: function (state) {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(state.list))
  },
  setItems: function (state, arr) {
    state.list = arr
  },
  loadItems: function (state) {
    state.list = JSON.parse(localStorage.getItem(STORAGE_KEY) || '[]')
  },
  addListItem: function (state, item) {
    state.list.push(item)
  },
  updateItem: function (state, { index, text }) {
    state.list[index].message = text
  },
  deleteItem: function (state, index) {
    state.list.splice(index, 1)
  },
}

export default {
  state,
  actions,
  mutations,
}