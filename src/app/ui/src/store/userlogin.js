import m from "mithril"; // eslint-disable-line no-unused-vars
import Submit from "@/module/submit";
import Flash from "@/component/flash";
import CookieStore from "@/module/cookiestore";

var UserLogin = (e, user) => {
  let login = () => {
    return m.request({
      method: "POST",
      url: "/api/v1/login",
      body: user,
    });
  };

  Submit.start(e);

  return login()
    .then((data) => {
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
      throw err;
    });
};

export default UserLogin;
