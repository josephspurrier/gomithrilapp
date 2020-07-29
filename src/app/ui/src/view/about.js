import m from "mithril"; // eslint-disable-line no-unused-vars
import SimplePage from "@/component/simple-page";

var Page = () => {
  return {
    view: () => (
      <SimplePage title="About">
        This shows you how to build a website using <strong>Mithril</strong>,{" "}
        <strong>Go</strong>, and <strong>Bulma</strong>.
      </SimplePage>
    ),
  };
};

export default Page;
