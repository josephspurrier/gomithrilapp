// eslint-disable-next-line no-unused-vars
import m from "mithril";
import {
  withKnobs,
  text,
  select,
  button,
  number,
  boolean,
} from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import Flash from "@/component/flash";
import "~/style/main.scss";

export default {
  title: "Component/Flash",
  component: Flash,
  decorators: [withKnobs, withA11y],
};

export const success = () => ({
  oninit: function () {
    Flash.timeout = -1;
    Flash.success(text("Text", "This is a success message."));
  },
  onremove: function () {
    Flash.clear();
  },
  view: () => <Flash />,
});

export const failed = () => ({
  oninit: function () {
    Flash.timeout = -1;
    Flash.failed(text("Text", "This is a failed message."));
  },
  onremove: function () {
    Flash.clear();
  },
  view: () => <Flash />,
});

export const warning = () => ({
  oninit: function () {
    Flash.timeout = -1;
    Flash.warning(text("Text", "This is a warning message."));
  },
  onremove: function () {
    Flash.clear();
  },
  view: () => <Flash />,
});

export const primary = () => ({
  oninit: function () {
    Flash.timeout = -1;
    Flash.primary(text("Text", "This is a primary message."));
  },
  onremove: function () {
    Flash.clear();
  },
  view: () => <Flash />,
});

export const link = () => ({
  oninit: function () {
    Flash.timeout = -1;
    Flash.link(text("Text", "This is a link message."));
  },
  onremove: function () {
    Flash.clear();
  },
  view: () => <Flash />,
});

export const info = () => ({
  oninit: function () {
    Flash.timeout = -1;
    Flash.info(text("Text", "This is a info message."));
  },
  onremove: function () {
    Flash.clear();
  },
  view: () => <Flash />,
});

export const dark = () => ({
  oninit: function () {
    Flash.timeout = -1;
    Flash.dark(text("Text", "This is a dark message."));
  },
  onremove: function () {
    Flash.clear();
  },
  view: () => <Flash />,
});

export const Action = () => ({
  oninit: function () {
    Flash.timeout = number("Timeout (milliseconds)", "2000");
    Flash.prepend = boolean("Prepend", false);
    let s = select(
      "Type",
      {
        success: "success",
        failed: "failed",
        warning: "warning",
        primary: "primary",
        link: "link",
        info: "info",
        dark: "dark",
      },
      "success"
    );
    Flash[s](text("Text", "This is a test message."));
    button("Show Message", function () {});
  },
  onremove: function () {
    Flash.clear();
  },
  view: () => <Flash />,
});
