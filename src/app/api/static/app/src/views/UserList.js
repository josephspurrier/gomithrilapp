// src/views/UserList.js
var m = require("mithril")
var User = require("../models/User")

module.exports = {
    oninit: User.loadList,
    view: function() {
        return m(".user-list", User.list.map(function(user) {
            return m(m.route.Link, {
                class: "user-list-item",
                href: "/edit/" + user.id,
            }, user.firstName + " " + user.lastName)
        }))
    }
}