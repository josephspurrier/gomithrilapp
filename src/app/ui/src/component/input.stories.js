import m from "mithril"; // eslint-disable-line no-unused-vars
import { withKnobs, text, select } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import Input from "./input";

export default {
  title: "Component/Input",
  component: Input,
  decorators: [withKnobs, withA11y],
};

export const input = () => ({
  oninit: (vnode) => {
    vnode.state.type = select(
      "Type",
      {
        text: "text",
        color: "color",
        date: "date",
        "datetime-local": "datetime-local",
        email: "email",
        hidden: "hidden",
        month: "month",
        number: "number",
        password: "password",
        range: "range",
        search: "search",
        time: "time",
        week: "week",
      },
      "text"
    );
  },
  view: (vnode) => (
    <Input
      label="First Name"
      value={text("Value", "John")}
      type={vnode.state.type}
    />
  ),
});
