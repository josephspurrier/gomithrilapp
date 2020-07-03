import m from "mithril";
import Auth from "@/module/auth";

function logout() {
  Auth.clear();
  m.route.set("/");
}

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
            <strong>gomithrilapp</strong>
          </m.route.Link>

          <a
            id="mobile-navbar-top"
            role="button"
            class="navbar-burger burger"
            aria-label="menu"
            aria-expanded="false"
            data-target="navbar-top"
            onclick={() => {
              const mob = document.getElementById("mobile-navbar-top");
              const nav = document.getElementById("navbar-top");
              mob.classList.toggle("is-active");
              nav.classList.toggle("is-active");
            }}
          >
            <span aria-hidden="true"></span>
            <span aria-hidden="true"></span>
            <span aria-hidden="true"></span>
          </a>
        </div>

        <div id="navbar-top" class="navbar-menu">
          <div class="navbar-end">
            <div class="navbar-item has-dropdown is-hoverable">
              <a class="navbar-link">Menu</a>

              <div class="navbar-dropdown is-right">
                {!Auth.isLoggedIn() && (
                  <m.route.Link class="navbar-item" href="/login">
                    Login
                  </m.route.Link>
                )}
                <a
                  class="navbar-item"
                  href={`https://petstore.swagger.io/?url=${location.origin}/static/swagger.json`}
                >
                  Swagger
                </a>
                <m.route.Link class="navbar-item" href="/about">
                  About
                </m.route.Link>
                <hr class="navbar-divider" />
                {Auth.isLoggedIn() && (
                  <a
                    class="dropdown-item"
                    onclick={() => {
                      logout();
                    }}
                  >
                    Logout
                  </a>
                )}
                <div class="navbar-item">v1.0.0</div>
              </div>
            </div>
          </div>
        </div>
      </nav>
    </main>
  ),
};

export default View;
