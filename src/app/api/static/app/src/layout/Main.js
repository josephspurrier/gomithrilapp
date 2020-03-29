var m = require('mithril')

var View = {
  view: function(vnode) {
    return m('main.layout', [
      m('nav.menu'),
      m('section', vnode.children)
    ])
  }
}

module.exports = View