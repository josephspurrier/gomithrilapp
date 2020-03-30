import m from "mithril";
import Submit from "../module/Submit";
import User from "../store/User";
import Flash from "./component/Flash";
import Auth from "../module/Auth";

var data = {
  title: "Login",
  subtitle: "Enter your login information below.",
};

function onsubmit(e) {
  Submit.start(e);

  User.login()
    .then((data) => {
      User.clear();
      Submit.finish();

      const auth = {
        accessToken: data.token,
        loggedIn: true,
      };

      Auth.save(auth);

      Flash.success("Login successful.");
      m.route.set("/");
    })
    .catch((err) => {
      Submit.finish();
      Flash.warning(err.response.message);
    });
}

var Page = {
  view: () => (
    <main>
      <div>
        <section class="section">
          <div class="container">
            <h1 class="title">{data.title}</h1>
            <h2 class="subtitle">{data.subtitle}</h2>
          </div>

          <div class="container" style="margin-top: 1em;">
            <form name="login" onsubmit={onsubmit}>
              <div class="field">
                <label class="label">Email</label>
                <div class="control">
                  <input
                    label="Email"
                    name="email"
                    type="text"
                    class="input"
                    data-cy="email"
                    required
                    oninput={(e) => {
                      User.current.email = e.target.value;
                    }}
                    value={User.current.email}
                  ></input>
                </div>
              </div>

              <div class="field">
                <label class="label">Password</label>
                <div class="control">
                  <input
                    label="Password"
                    name="password"
                    type="password"
                    class="input"
                    data-cy="password"
                    required
                    oninput={(e) => {
                      User.current.password = e.target.value;
                    }}
                    value={User.current.password}
                  ></input>
                </div>
              </div>

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
                      User.clear();
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

export default Page;
