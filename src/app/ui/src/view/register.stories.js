import m from "mithril"; // eslint-disable-line no-unused-vars
import { withKnobs, select, text } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import RegisterPage from "@/view/register";
import Flash from "@/component/flash";
import MockRequest from "@/component/mockrequest";
import "~/style/main.scss";

export default {
  title: "View/Register",
  component: RegisterPage,
  decorators: [withKnobs, withA11y],
};

export const register = () => ({
  oninit: () => {
    let s = select(
      "Operation",
      {
        UserRegistered: "opt1",
        UserAlreadyExists: "opt2",
      },
      "opt1"
    );
    switch (s) {
      case "opt1":
        MockRequest.ok({});
        break;
      case "opt2":
        MockRequest.badRequest("The user already exists.");
        break;
      default:
        MockRequest.badRequest("There is a problem with the storybook.");
    }
  },
  view: () => (
    <main>
      <RegisterPage
        firstName={text("First Name", "Joe")}
        lastName={text("Last Name", "Smith")}
        email={text("Email", "jsmith@example.com")}
        password={text("Password", "password")}
      />
      <Flash />
    </main>
  ),
});
