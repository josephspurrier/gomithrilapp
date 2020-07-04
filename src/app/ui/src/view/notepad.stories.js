// eslint-disable-next-line no-unused-vars
import m from "mithril";
import { withKnobs } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import NotepadPage from "@/view/notepad";
import "~/style/main.scss";

export default {
  title: "View/Notepad",
  component: NotepadPage,
  decorators: [withKnobs, withA11y],
};

export const notepad = () => ({
  view: () => <NotepadPage />,
});
