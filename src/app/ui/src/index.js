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
      if (!CookieStore.isLoggedIn()) m.route.set("/login");
    },
    render: () => m(LayoutMain, m(HomePage)),
  },
  "/notepad": {
    onmatch: () => {
      if (!CookieStore.isLoggedIn()) m.route.set("/login");
    },
    render: () => m(LayoutMain, m(NotepadPage)),
  },
  "/login": {
    onmatch: () => {
      if (CookieStore.isLoggedIn()) m.route.set("/");
    },
    render: () => m(LayoutMain, m(LoginPage)),
  },
  "/register": {
    onmatch: () => {
      if (CookieStore.isLoggedIn()) m.route.set("/");
    },
    render: () => m(LayoutMain, m(RegisterPage)),
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
