import m from "mithril";
import Submit from "@/module/submit";
import Flash from "@/component/flash";

var UserRegister = {
  user: {
    first_name: "",
    last_name: "",
    email: "",
    password: "",
  },
  clear: () => {
    UserRegister.user = {};
  },
  register: () => {
    return m.request({
      method: "POST",
      url: "/api/v1/register",
      body: UserRegister.user,
    });
  },
  submit: function (e) {
    Submit.start(e);

    UserRegister.register()
      .then(() => {
        UserRegister.clear();
        Submit.finish();

        Flash.success("User registered.");
        m.route.set("/login");
      })
      .catch((err) => {
        Submit.finish();
        Flash.warning(err.response.message);
      });
  },
};

export default UserRegister;
