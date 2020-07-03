import m from "mithril";

var data = {
  title: "Welcome",
  subtitle: "Login was successful",
};

var Page = {
  view: () =>
    m(
      "home",
      <div>
        <section class="hero is-primary">
          <div class="hero-body">
            <div class="container">
              <h1 class="title">{data.title}</h1>
              <h2 class="subtitle">{data.subtitle}</h2>
            </div>
          </div>
        </section>
        <br />
        <div class="container">
          <m.route.Link href="/notepad" data-cy="notepad-link">
            Click here to access your Notepad.
          </m.route.Link>
        </div>
      </div>
    ),
};

export default Page;
