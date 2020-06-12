// eslint-disable-next-line no-unused-vars
import m from "mithril";
import { action } from "@storybook/addon-actions";
import { withKnobs, text, boolean, number } from "@storybook/addon-knobs";
import Block from "./module/block";

export default {
  title: "Structure/Block",
  component: Block,
  decorators: [withKnobs],
};

export const button = () => ({
  view: () => (
    <button
      disabled={boolean("Disabled", false)}
      onclick={action("button-click")}
    >
      {text("Label", "Hello Storybook")}
    </button>
  ),
});

export const DynamicText = () => ({
  view: () => {
    const name = text("Name", "Joe");
    const age = number("Age", 32);
    const content = `I am ${name} and I'm ${age} years old.`;

    return m("", content);
  },
});

export const long = () => {
  return {
    view: () => <Block>Long</Block>,
  };
};
long.story = {
  name: "Long",
};

export const short = () => ({
  view: () => <Block>Short</Block>,
});

export const emoji = () => ({
  view: () => (
    <Block>
      <form>
        <span role="img" aria-label="so cool">
          😀 😎 👍 💯
        </span>
      </form>
    </Block>
  ),
});
