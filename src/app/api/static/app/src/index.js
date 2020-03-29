import m, { route } from "mithril";
import UserList from "./page/UserList";
import UserForm from "./page/UserForm";
import AboutPage from "./page/About";
import LoginPage from "./page/Login";
import RegisterPage from "./page/Register";
import LayoutMain from "./layout/Main";

route(document.body, "/list", {
  "/list": {
    render: function () {
      return m(LayoutMain, m(UserList));
    },
  },
  "/login": {
    render: function () {
      return m(LayoutMain, m(LoginPage));
    },
  },
  "/register": {
    render: function () {
      return m(LayoutMain, m(RegisterPage));
    },
  },
  "/about": {
    render: function () {
      return m(LayoutMain, m(AboutPage));
    },
  },
  "/edit/:id": {
    render: function (vnode) {
      return m(LayoutMain, m(UserForm, vnode.attrs));
    },
  },
});
