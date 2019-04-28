import shortid from 'shortid'

export default {
  STORAGE_KEY: 'todos-vuejs-2.0',
  list: [],
  saveItems: function() {
    localStorage.setItem(this.STORAGE_KEY, JSON.stringify(this.list))
  },
  loadItems: function() {
    this.list = JSON.parse(localStorage.getItem(this.STORAGE_KEY) || '[]')
    // console.log('Loaded items:', this.list.length)
  },
  addItem: function() {
    const inputTodo = document.getElementById('inputTodo')
    this.list.push({ id: shortid.generate(), message: inputTodo.value })
    inputTodo.value = ''
    this.saveItems()
  },
  deleteItem: function(position) {
    this.list.splice(position, 1)
    this.saveItems()
  },
  updateItem: function(position, text) {
    this.list[position].message = text
    this.saveItems()
  }
}
