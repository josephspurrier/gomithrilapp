import m from "mithril"; // eslint-disable-line no-unused-vars
import { withKnobs, text } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import LayoutMain from "./main";
import SimplePage from "@/component/simple-page";
import "~/style/main.scss";

export default {
  title: "Component/Layout Main",
  component: LayoutMain,
  decorators: [withKnobs, withA11y],
};

export const simplePage = () => ({
  view: () => {
    return m(
      LayoutMain,
      m(
        SimplePage,
        {
          title: text("Title", "This is the Title"),
          description: text(
            "Description",
            "This is a subtitle or description."
          ),
        },
        [text("Content", "This is the content.")]
      )
    );
  },
});
