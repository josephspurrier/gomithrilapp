import m from "mithril"; // eslint-disable-line no-unused-vars
import SimplePage from "@/component/simple-page";

var Page = () => {
  return {
    view: () => (
      <SimplePage title="Error" description="The page is not found." />
    ),
  };
};

export default Page;
