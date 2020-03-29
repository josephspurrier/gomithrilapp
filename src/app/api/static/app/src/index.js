import m, { route } from "mithril";
import AboutPage from "./page/About";
import LoginPage from "./page/Login";
import RegisterPage from "./page/Register";
import IndexPage from "./page/Index";
import NotepadPage from "./page/Notepad";
import LayoutMain from "./layout/Main";
import Cookie from "js-cookie";

function isLoggedIn() {
  try {
    let auth = Cookie.get("auth");
    if (auth === undefined) {
      return false;
    }
    return true;
  } catch (err) {
    console.log(err);
  }

  return false;
}

route(document.body, "/", {
  "/": {
    onmatch: function () {
      if (isLoggedIn()) return Index;
      else m.route.set("/login");
    },
  },
  "/notepad": {
    onmatch: function () {
      if (isLoggedIn()) return Notepad;
      else m.route.set("/login");
    },
  },
  "/login": {
    onmatch: function () {
      if (isLoggedIn()) m.route.set("/");
      else return Login;
    },
  },
  "/register": {
    onmatch: function () {
      if (isLoggedIn()) m.route.set("/");
      else return Register;
    },
  },
  "/about": {
    render: function () {
      return m(LayoutMain, m(AboutPage));
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
