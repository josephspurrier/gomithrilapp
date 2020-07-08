import m from "mithril"; // eslint-disable-line no-unused-vars
import { withKnobs, select, text } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import LoginPage from "@/view/login";
import Flash from "@/component/flash";
import MockRequest from "@/module/mockrequest";
import "~/style/main.scss";

export default {
  title: "View/Login",
  component: LoginPage,
  decorators: [withKnobs, withA11y],
};

export const login = () => ({
  oninit: () => {
    let s = select(
      "Operation",
      {
        LoginSuccessful: "opt1",
        LoginIncorrect: "opt2",
      },
      "opt1"
    );
    switch (s) {
      case "opt1":
        MockRequest.ok({});
        break;
      case "opt2":
        MockRequest.badRequest("Login information does not match.");
        break;
      default:
        MockRequest.badRequest("There is a problem with the storybook.");
    }
  },
  view: () => (
    <main>
      <LoginPage
        email={text("Email", "jsmith@example.com")}
        password={text("Password", "password")}
      />
      <Flash />
    </main>
  ),
});
