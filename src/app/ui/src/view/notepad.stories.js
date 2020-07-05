// eslint-disable-next-line no-unused-vars
import m from "mithril";
import { withKnobs } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import NotepadPage from "@/view/notepad";
import FlashContainer from "@/component/flashcontainer";
import Mock from "@/component/mock";
import "~/node_modules/@fortawesome/fontawesome-free/js/all.js";
import "~/style/main.scss";

export default {
  title: "View/Notepad",
  component: NotepadPage,
  decorators: [withKnobs, withA11y],
};

export const notepad = () => ({
  oninit: () => {
    Mock.success({
      notes: [
        {
          id: "6e8568e5-2632-7c8d-b448-ec82772ed4ec",
          message: "foo",
        },
        {
          id: "a3969708-bf1c-efd4-9d98-d8d5a217cd93",
          message: "bar",
        },
      ],
    });
  },
  view: () => (
    <main>
      <NotepadPage />
      <FlashContainer />
    </main>
  ),
});
