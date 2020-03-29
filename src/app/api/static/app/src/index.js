var m = require('mithril')

var UserList = require('./view/UserList')
var UserForm = require('./view/UserForm')
var LayoutMain = require('./layout/Main')
var LayoutList = require('./layout/List')

m.route(document.body, '/list', {
  '/list': {
    render: function() {
      return m(LayoutMain, m(UserList))
    }
  },
  '/edit/:id': {
    render: function(vnode) {
      return m(LayoutList, m(UserForm, vnode.attrs))
    }
  },
})