import m from "mithril"; // eslint-disable-line no-unused-vars

var Page = () => {
  return {
    view: ({ attrs, children }) =>
      m(
        "page",
        <div>
          <section class="section">
            <div class="container">
              <h1 class="title">{attrs.title}</h1>
              <h2 class="subtitle">{attrs.description}</h2>
              {children}
            </div>
          </section>
        </div>
      ),
  };
};

export default Page;
