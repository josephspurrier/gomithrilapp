import m from "mithril"; // eslint-disable-line no-unused-vars
import { withKnobs, boolean, text } from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import LoginPage from "@/view/login";
import Flash from "@/component/flash";
import { rest } from "msw";
import { worker } from "@/mock/browser";

export default {
  title: "View/Login",
  component: LoginPage,
  decorators: [withKnobs, withA11y],
};

export const login = () => ({
  oninit: () => {
    const shouldFail = boolean("Fail", false);

    worker.use(
      ...[
        rest.post("/api/v1/login", (req, res, ctx) => {
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
    <main>
      <LoginPage
        email={text("Email", "jsmith@example.com")}
        password={text("Password", "password")}
      />
      <Flash />
    </main>
  ),
});
