import m from "mithril";

var Page = {
  view: (vnode) =>
    m(
      "page",
      <div>
        <section class="section">
          <div class="container">
            <h1 class="title">{vnode.attrs.title}</h1>
            <h2 class="subtitle">{vnode.attrs.description}</h2>
            <div>{vnode.attrs.content}</div>
          </div>
        </section>
      </div>
    ),
};

export default Page;