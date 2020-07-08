import m from "mithril"; // eslint-disable-line no-unused-vars
import { withKnobs, select } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import Note from "./note";
import Flash from "@/component/flash";
import MockRequest from "@/module/mockrequest";
import "~/style/main.scss";

export default {
  title: "Component/Note",
  component: Note,
  decorators: [withKnobs, withA11y],
};

export const note = () => ({
  oninit: () => {
    let s = select(
      "Operation",
      {
        Success: "opt1",
        Fail: "opt2",
      },
      "opt1"
    );
    switch (s) {
      case "opt1":
        MockRequest.ok({});
        break;
      case "opt2":
        MockRequest.badRequest("There was an error.");
        break;
      default:
        MockRequest.badRequest("There is a problem with the storybook.");
    }
  },
  view: () => (
    <app>
      <ul>
        <Note />
      </ul>
      <Flash />
    </app>
  ),
});
