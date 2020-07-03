import m from "mithril";

var data = {
  title: "Error",
};

var Page = {
  view: () =>
    m(
      "error",
      <div>
        <section class="section">
          <div class="container">
            <h1 class="title">{data.title}</h1>
            <h2 class="subtitle">The page is not found.</h2>
          </div>
        </section>
      </div>
    ),
};

export default Page;
