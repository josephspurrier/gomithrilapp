import m from "mithril"; // eslint-disable-line no-unused-vars
import { withKnobs, text } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import SimplePage from "@/component/simple-page";
import "~/style/main.scss";

export default {
  title: "Component/Simple Page",
  component: SimplePage,
  decorators: [withKnobs, withA11y],
};

export const populated = () => ({
  view: () => (
    <SimplePage
      title={text("Title", "This is the Title")}
      description={text("Description", "This is a subtitle or description.")}
    >
      {text("Content", "This is the content.")}
    </SimplePage>
  ),
});
populated.story = {
  name: "With Content",
};
