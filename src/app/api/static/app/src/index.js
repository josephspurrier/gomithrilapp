var m = require('mithril')

var UserList = require('./page/UserList')
var UserForm = require('./page/UserForm')
var AboutPage = require('./page/About')
var LoginPage = require('./page/Login')
var RegisterPage = require('./page/Register')
var LayoutMain = require('./layout/Main')

m.route(document.body, '/list', {
  '/list': {
    render: function() {
      return m(LayoutMain, m(UserList))
    }
  },
  '/login': {
    render: function() {
      return m(LayoutMain, m(LoginPage))
    }
  },
  '/register': {
    render: function() {
      return m(LayoutMain, m(RegisterPage))
    }
  },
  '/about': {
    render: function() {
      return m(LayoutMain, m(AboutPage))
    }
  },
  '/edit/:id': {
    render: function(vnode) {
      return m(LayoutMain, m(UserForm, vnode.attrs))
    }
  },
})