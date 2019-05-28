import shortid from 'shortid'

const STORAGE_KEY = 'app-note'

export const state = () => ({
  list: [],
})

export const actions = {
  addItem: function (state, value) {
    state.commit('addListItem', { id: shortid.generate(), message: value })
    state.commit('save')
  },
  updateItem: function (state, { id, text }) {
    state.commit('updateItem', { id, text })
    state.commit('save')
  },
  deleteItem: function (state, id) {
    state.commit('deleteItem', id)
    state.commit('save')
  },
  loadItems: function (state) {
    state.commit('loadItems')
  },
}

export const mutations = {
  save: function (state) {
    localStorage.setItem(STORAGE_KEY, JSON.stringify(state.list))
  },
  loadItems: function (state) {
    state.list = JSON.parse(localStorage.getItem(STORAGE_KEY) || '[]')
  },
  addListItem: function (state, item) {
    state.list.push(item)
  },
  updateItem: function (state, { id, text }) {
    state.list[id].message = text
  },
  deleteItem: function (state, id) {
    state.list.splice(id, 1)
  },
}

export default {
  state,
  actions,
  mutations,
}