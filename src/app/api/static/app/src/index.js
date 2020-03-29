import m, { route } from "mithril";
import UserList from "./page/UserList";
import UserForm from "./page/UserForm";
import AboutPage from "./page/About";
import LoginPage from "./page/Login";
import RegisterPage from "./page/Register";
import LayoutMain from "./layout/Main";
import Cookie from "js-cookie";

function isLoggedIn() {
  try {
    let auth = Cookie.get("auth");
    if (auth === undefined) {
      return false;
    }

    console.log("Cookie:", auth);
    let v = JSON.parse(auth);

    console.log("Auth:", v);
    return true;
  } catch (err) {
    console.log(err);
  }

  return false;
}

route(document.body, "/", {
  "/": {
    onmatch: function () {
      if (isLoggedIn()) m.route.set("/list");
      else m.route.set("/login");
    },
  },
  "/list": {
    onmatch: function () {
      if (isLoggedIn()) return List;
      else m.route.set("/login");
    },
  },
  "/edit/:id": {
    onmatch: function () {
      if (isLoggedIn()) return Edit;
      else m.route.set("/login");
    },
  },
  "/login": {
    onmatch: function () {
      if (isLoggedIn()) m.route.set("/list");
      else return Login;
    },
  },
  "/register": {
    onmatch: function () {
      if (isLoggedIn()) m.route.set("/list");
      else return Register;
    },
  },
  "/about": {
    render: function () {
      return m(LayoutMain, m(AboutPage));
    },
  },
});

var List = {
  view: function () {
    return m(LayoutMain, m(UserList));
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

var Edit = {
  view: function (vnode) {
    return m(LayoutMain, m(UserForm, vnode.attrs));
  },
};
