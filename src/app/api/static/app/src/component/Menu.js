import m from "mithril";

var View = {
  view: () => (
    <main>
      <nav
        class="navbar is-black"
        role="navigation"
        aria-label="main navigation"
      >
        <div class="navbar-brand">
          <m.route.Link class="navbar-item" href="/" data-cy="home-link">
            <strong>govueapp</strong>
          </m.route.Link>

          <button class="button navbar-burger">
            <span />
            <span />
            <span />
          </button>
        </div>

        <div class="navbar-end">
          <div class="navbar-item has-dropdown is-hoverable">
            <a class="navbar-link">Menu</a>

            <div class="navbar-dropdown is-right">
              <m.route.Link class="navbar-item" href="/login">
                Login
              </m.route.Link>
              <m.route.Link class="navbar-item" href="/about">
                About
              </m.route.Link>
              <hr class="navbar-divider" />
              <a v-if="isAuthenticated" class="dropdown-item">
                Logout
              </a>
              <div class="navbar-item">v1.0.0</div>
            </div>
          </div>
        </div>
      </nav>
    </main>
  ),
};

export default View;
