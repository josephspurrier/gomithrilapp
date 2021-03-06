import m from "mithril"; // eslint-disable-line no-unused-vars
import UserLogin from "@/store/userlogin";
import Input from "@/component/input";

var Page = ({ attrs }) => {
  var data = {
    title: "Login",
    subtitle: "Enter your login information below.",
  };

  let user = {};

  var clear = () => {
    user = {};
  };

  // Prefill the fields.
  if (attrs.email) {
    user.email = attrs.email;
  }
  if (attrs.password) {
    user.password = attrs.password;
  }

  return {
    view: (vnode) => (
      <main>
        <div>
          <section class="section">
            <div class="container">
              <h1 class="title">{data.title}</h1>
              <h2 class="subtitle">{data.subtitle}</h2>
            </div>

            <div class="container" style="margin-top: 1em;">
              <form
                name="login"
                onsubmit={(e) => {
                  UserLogin(e, user)
                    .then(() => {
                      clear(vnode);
                    })
                    .catch(() => {});
                }}
              >
                <Input
                  label="Email"
                  name="email"
                  required="true"
                  oninput={(e) => {
                    user.email = e.target.value;
                  }}
                  value={user.email}
                />

                <Input
                  label="Password"
                  name="password"
                  required="true"
                  oninput={(e) => {
                    user.password = e.target.value;
                  }}
                  value={user.password}
                  type="password"
                />

                <div class="field is-grouped">
                  <p class="control">
                    <button
                      id="submit"
                      type="submit"
                      data-cy="submit"
                      class="button is-primary"
                    >
                      Submit
                    </button>
                  </p>

                  <p class="control">
                    <button
                      type="button"
                      class="button is-light"
                      onclick={() => {
                        clear(vnode);
                      }}
                    >
                      Clear
                    </button>
                  </p>

                  <p class="control">
                    <m.route.Link href="/register" class="button is-light">
                      Register
                    </m.route.Link>
                  </p>
                </div>
              </form>
            </div>
          </section>
        </div>
      </main>
    ),
  };
};

export default Page;
