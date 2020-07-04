// eslint-disable-next-line no-unused-vars
import m from "mithril";
import {
  withKnobs,
  text,
  select,
  button,
  number,
} from "@storybook/addon-knobs";
import { withA11y } from "@storybook/addon-a11y";
import Flash from "@/component/flash";
import FlashContainer from "@/component/flashcontainer";
import "~/style/main.scss";

export default {
  title: "Component/Flash",
  component: Flash,
  decorators: [withKnobs, withA11y],
};

export const success = () => ({
  oncreate: function () {
    Flash.timeout = -1;
    Flash.success(text("Text", "This is a success message."));
  },
  view: () => <FlashContainer />,
});

export const failed = () => ({
  oncreate: function () {
    Flash.timeout = -1;
    Flash.failed(text("Text", "This is a failed message."));
  },
  view: () => <FlashContainer />,
});

export const warning = () => ({
  oncreate: function () {
    Flash.timeout = -1;
    Flash.warning(text("Text", "This is a warning message."));
  },
  view: () => <FlashContainer />,
});

export const primary = () => ({
  oncreate: function () {
    Flash.timeout = -1;
    Flash.primary(text("Text", "This is a primary message."));
  },
  view: () => <FlashContainer />,
});

export const link = () => ({
  oncreate: function () {
    Flash.timeout = -1;
    Flash.link(text("Text", "This is a link message."));
  },
  view: () => <FlashContainer />,
});

export const info = () => ({
  oncreate: function () {
    Flash.timeout = -1;
    Flash.info(text("Text", "This is a info message."));
  },
  view: () => <FlashContainer />,
});

export const dark = () => ({
  oncreate: function () {
    Flash.timeout = -1;
    Flash.dark(text("Text", "This is a dark message."));
  },
  view: () => <FlashContainer />,
});

export const Action = () => ({
  oncreate: function () {
    Flash.timeout = number("Timeout (milliseconds)", "2000");
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
  view: () => <FlashContainer />,
});
