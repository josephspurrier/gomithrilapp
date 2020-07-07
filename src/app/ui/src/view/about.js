import m from "mithril"; // eslint-disable-line no-unused-vars
import SimplePage from "@/component/simple-page";

var data = {
  title: "About",
  description: {
    view: () =>
      m(
        "description",
        <main>
          This shows you how to build a website using <strong>Mithril</strong>,{" "}
          <strong>Go</strong>, and <strong>Bulma </strong>
        </main>
      ),
  },
};

var Page = {
  view: () => (
    <SimplePage title={data.title} description={m(data.description)} />
  ),
};

export default Page;
