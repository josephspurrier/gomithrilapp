var m = require('mithril')

var User = {
  current: {},
  list: [],
  loadList: function() {
    return m.request({
      method: 'GET',
      url: 'https://rem-rest-api.herokuapp.com/api/users',
      withCredentials: true,
    })
    .then(function(result){
      User.list = result.data
    })
  },
  load: function(id) {
    return m.request({
      method: 'GET',
      url: 'https://rem-rest-api.herokuapp.com/api/users/'+id,
      withCredentials: true,
    })
    .then(function(result){
      User.current = result
    })
  },
  save: function() {
    return m.request({
      method: 'PUT',
      url: 'https://rem-rest-api.herokuapp.com/api/users/'+User.current.id,
      withCredentials: true,
      body: User.current
    })
  }
}

module.exports = User