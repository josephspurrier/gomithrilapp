var m = require('mithril')
var User = require('../store/User')
var Submit = require('../module/Submit')
var Sleep = require('../module/Sleep')

var View = {
  oninit: function(vnode){
    User.load(vnode.attrs.id)
  },
  view: function() {
    return m('form', {
      onsubmit: function(e){
        Submit.start(e)

        Sleep(500).then(() => {
          User.save().then(() => {
            m.route.set('/list')
          }).catch(function (e){
            alert('Could not save content.',e)
          }).finally(function() {
            Submit.finish()
          })
      })
      }
    }, [
      m('label.label', 'First Name'),
      m('input.input[type=text][placeholder=First name]',{
        oninput: function(e) {
          User.current.firstName = e.target.value
        },
        value: User.current.firstName,
      }),
      m('label.label', 'Last name'),
      m('input.input[placeholder=Last name]', {
        oninput: function(e) {
          User.current.lastName = e.target.value
        },
        value: User.current.lastName,
      }),
      m('button.button[type=submit]',{
        disabled: Submit.disabled,
      },Submit.text('Save'))
    ])
  }
}

module.exports = View