import m from "mithril";
import AboutPage from "./page/about";
import LoginPage from "./page/login";
import RegisterPage from "./page/register";
import IndexPage from "./page/index";
import NotepadPage from "./page/notepad";
import ErrorPage from "./page/error";
import LayoutMain from "./page/layout/main";
import Auth from "./module/auth";
import "../node_modules/@fortawesome/fontawesome-free/js/all.js";
import "../style/main.scss";

m.route.prefix = "";

m.route(document.body, "/", {
  "/": {
    onmatch: function () {
      if (Auth.isLoggedIn()) return Index;
      else m.route.set("/login");
    },
  },
  "/notepad": {
    onmatch: function () {
      if (Auth.isLoggedIn()) return Notepad;
      else m.route.set("/login");
    },
  },
  "/login": {
    onmatch: function () {
      if (Auth.isLoggedIn()) m.route.set("/");
      else return Login;
    },
  },
  "/register": {
    onmatch: function () {
      if (Auth.isLoggedIn()) m.route.set("/");
      else return Register;
    },
  },
  "/about": {
    render: function () {
      return m(LayoutMain, m(AboutPage));
    },
  },
  "/:404...": {
    view: function () {
      return m(LayoutMain, m(ErrorPage));
    },
  },
});

var Index = {
  view: function () {
    return m(LayoutMain, m(IndexPage));
  },
};

var Notepad = {
  view: function () {
    return m(LayoutMain, m(NotepadPage));
  },
};

var Login = {
  view: function () {
    return m(LayoutMain, m(LoginPage));
  },
};

var Register = {
  view: function () {
    return m(LayoutMain, m(RegisterPage));
  },
};
