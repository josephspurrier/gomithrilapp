var m = require('mithril')
var Menu = require('../component/Menu')

var View = {
  view: function(vnode) {
    return m('main.layout', [
      m(Menu),
      m('section', vnode.children)
    ])
  }
}

module.exports = View