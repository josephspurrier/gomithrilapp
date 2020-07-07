import m from "mithril"; // eslint-disable-line no-unused-vars
import { withKnobs } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import HomePage from "@/view/home";
import "~/style/main.scss";

export default {
  title: "View/Home",
  component: HomePage,
  decorators: [withKnobs, withA11y],
};

export const home = () => ({
  view: () => <HomePage />,
});
