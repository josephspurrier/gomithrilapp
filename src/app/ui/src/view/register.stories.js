// eslint-disable-next-line no-unused-vars
import m from "mithril";
import { withKnobs } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import RegisterPage from "@/view/register";
import "~/style/main.scss";

export default {
  title: "View/Register",
  component: RegisterPage,
  decorators: [withKnobs, withA11y],
};

export const register = () => ({
  view: () => <RegisterPage />,
});
