import m from "mithril";
import UserLogin from "@/store/userlogin";

var data = {
  title: "Login",
  subtitle: "Enter your login information below.",
};

var Page = {
  onremove: () => {
    UserLogin.clear();
  },
  view: () => (
    <main>
      <div>
        <section class="section">
          <div class="container">
            <h1 class="title">{data.title}</h1>
            <h2 class="subtitle">{data.subtitle}</h2>
          </div>

          <div class="container" style="margin-top: 1em;">
            <form name="login" onsubmit={UserLogin.submit}>
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
                      UserLogin.user.email = e.target.value;
                    }}
                    value={UserLogin.user.email}
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
                      UserLogin.user.password = e.target.value;
                    }}
                    value={UserLogin.user.password}
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
                      UserLogin.clear();
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
