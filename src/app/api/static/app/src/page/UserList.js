var m = require('mithril')
var User = require('../store/User')

var View = {
  oninit: User.loadList,
  view: function() {
    return m('.user-list', User.list.map(function(user){
      return m(m.route.Link,{
        class: 'user-list-item',
        href: '/edit/'+ user.id,
      }, user.firstName + ' ' + user.lastName)
    }))
  }
}

module.exports = View