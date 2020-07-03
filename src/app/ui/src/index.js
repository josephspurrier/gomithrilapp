import m from "mithril";
import AboutPage from "@/page/about";
import LoginPage from "@/view/login";
import RegisterPage from "@/view/register";
import HomePage from "@/page/home";
import NotepadPage from "@/page/notepad";
import ErrorPage from "@/page/error";
import LayoutMain from "@/page/layout/main";
import CookieStore from "@/module/cookiestore";
import "~/node_modules/@fortawesome/fontawesome-free/js/all.js";
import "~/style/main.scss";

m.route.prefix = "";

m.route(document.body, "/", {
  "/": {
    onmatch: function () {
      if (CookieStore.isLoggedIn()) return Index;
      else m.route.set("/login");
    },
  },
  "/notepad": {
    onmatch: function () {
      if (CookieStore.isLoggedIn()) return Notepad;
      else m.route.set("/login");
    },
  },
  "/login": {
    onmatch: function () {
      if (CookieStore.isLoggedIn()) m.route.set("/");
      else return Login;
    },
  },
  "/register": {
    onmatch: function () {
      if (CookieStore.isLoggedIn()) m.route.set("/");
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
    return m(LayoutMain, m(HomePage));
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
