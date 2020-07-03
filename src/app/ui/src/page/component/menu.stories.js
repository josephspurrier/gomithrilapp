// eslint-disable-next-line no-unused-vars
import m from "mithril";
//import { action } from "@storybook/addon-actions";
import { withKnobs } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import Menu from "./menu";
import "../../../style/main.scss";

export default {
  title: "Structure/Menu",
  component: Menu,
  decorators: [withKnobs, withA11y],
};

export const menu = () => ({
  view: () => <Menu />,
});
menu.story = {
  name: "Menu",
};
