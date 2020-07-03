import m from "mithril";
import Submit from "@/module/submit";
import Flash from "@/page/component/flash";
import CookieStore from "~/src/module/cookiestore";

var UserLogin = {
  user: {
    email: "",
    password: "",
  },
  clear: () => {
    UserLogin.user = {};
  },
  login: () => {
    return m.request({
      method: "POST",
      url: "/api/v1/login",
      body: UserLogin.user,
    });
  },
  onSubmit: function (e) {
    Submit.start(e);

    UserLogin.login()
      .then((data) => {
        UserLogin.clear();
        Submit.finish();

        const auth = {
          accessToken: data.token,
          loggedIn: true,
        };

        CookieStore.save(auth);

        Flash.success("Login successful.");
        m.route.set("/");
      })
      .catch((err) => {
        Submit.finish();
        Flash.warning(err.response.message);
      });
  },
};

export default UserLogin;
