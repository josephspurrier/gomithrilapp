import m from "mithril"; // eslint-disable-line no-unused-vars
import AboutPage from "@/view/about";
import LoginPage from "@/view/login";
import RegisterPage from "@/view/register";
import HomePage from "@/view/home";
import NotepadPage from "@/view/notepad";
import ErrorPage from "@/view/error";
import LayoutMain from "@/layout/main";
import CookieStore from "@/module/cookiestore";
import "~/node_modules/@fortawesome/fontawesome-free/js/all.js";
import "~/style/main.scss";

m.route.prefix = "";

m.route(document.body, "/", {
  "/": {
    onmatch: () => {
      if (CookieStore.isLoggedIn()) return Index;
      else m.route.set("/login");
    },
  },
  "/notepad": {
    onmatch: () => {
      if (CookieStore.isLoggedIn()) return Notepad;
      else m.route.set("/login");
    },
  },
  "/login": {
    onmatch: () => {
      if (CookieStore.isLoggedIn()) m.route.set("/");
      else return Login;
    },
  },
  "/register": {
    onmatch: () => {
      if (CookieStore.isLoggedIn()) m.route.set("/");
      else return Register;
    },
  },
  "/about": {
    render: () => {
      return m(LayoutMain, m(AboutPage));
    },
  },
  "/:404...": {
    view: () => {
      return m(LayoutMain, m(ErrorPage));
    },
  },
});

var Index = {
  view: () => {
    return m(LayoutMain, m(HomePage));
  },
};

var Notepad = {
  view: () => {
    return m(LayoutMain, m(NotepadPage));
  },
};

var Login = {
  view: () => {
    return m(LayoutMain, m(LoginPage));
  },
};

var Register = {
  view: () => {
    return m(LayoutMain, m(RegisterPage));
  },
};
