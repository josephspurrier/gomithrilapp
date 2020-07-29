import m from "mithril"; // eslint-disable-line no-unused-vars
import { withKnobs, boolean } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import Note from "./note";
import Flash from "@/component/flash";
import { rest } from "msw";
import { worker } from "@/mock/browser";

export default {
  title: "Component/Note",
  component: Note,
  decorators: [withKnobs, withA11y],
};

export const note = () => ({
  oninit: () => {
    const shouldFail = boolean("Fail", false);

    worker.use(
      ...[
        rest.put("/api/v1/note/1", (req, res, ctx) => {
          if (shouldFail) {
            return res(
              ctx.status(400),
              ctx.json({
                message: "There was an error.",
              })
            );
          } else {
            return res(
              ctx.status(200),
              ctx.json({
                message: "ok",
              })
            );
          }
        }),
        rest.delete("/api/v1/note/1", (req, res, ctx) => {
          if (shouldFail) {
            return res(
              ctx.status(400),
              ctx.json({
                message: "There was an error.",
              })
            );
          } else {
            return res(
              ctx.status(200),
              ctx.json({
                message: "ok",
              })
            );
          }
        }),
      ]
    );
  },
  view: () => (
    <app>
      <ul>
        <Note id="1" />
      </ul>
      <Flash />
    </app>
  ),
});
