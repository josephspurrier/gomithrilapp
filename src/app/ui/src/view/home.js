import m from "mithril"; // eslint-disable-line no-unused-vars

var Page = () => {
  return {
    view: () =>
      m(
        "home",
        <div>
          <section class="hero is-primary">
            <div class="hero-body">
              <div class="container">
                <h1 class="title">Welcome</h1>
                <h2 class="subtitle">Login was successful.</h2>
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
};

export default Page;
