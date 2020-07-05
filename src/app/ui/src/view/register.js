// eslint-disable-next-line no-unused-vars
import m from "mithril";
import UserRegister from "@/store/userregister";

var data = {
  title: "Register",
  subtitle: "Enter your information below.",
};

var Page = {
  onremove: () => {
    UserRegister.clear();
  },
  oncreate: (vnode) => {
    // Prefill the fields.
    if (vnode.attrs.firstName) {
      UserRegister.user.first_name = vnode.attrs.firstName;
    }
    if (vnode.attrs.lastName) {
      UserRegister.user.last_name = vnode.attrs.lastName;
    }
    if (vnode.attrs.email) {
      UserRegister.user.email = vnode.attrs.email;
    }
    if (vnode.attrs.password) {
      UserRegister.user.password = vnode.attrs.password;
    }

    m.redraw();
  },
  view: () => {
    return m(
      "",
      <main>
        <div>
          <section class="section">
            <div class="container">
              <h1 class="title">{data.title}</h1>
              <h2 class="subtitle">{data.subtitle}</h2>
            </div>

            <div class="container" style="margin-top: 1em;">
              <form name="login" onsubmit={UserRegister.submit}>
                <div class="field">
                  <label class="label">First Name</label>
                  <div class="control">
                    <input
                      label="first_name"
                      name="first_name"
                      type="text"
                      class="input"
                      data-cy="first_name"
                      required
                      oninput={(e) => {
                        UserRegister.user.first_name = e.target.value;
                      }}
                      value={UserRegister.user.first_name}
                    ></input>
                  </div>
                </div>

                <div class="field">
                  <label class="label">Last Name</label>
                  <div class="control">
                    <input
                      label="last_name"
                      name="last_name"
                      type="text"
                      class="input"
                      data-cy="last_name"
                      required
                      oninput={(e) => {
                        UserRegister.user.last_name = e.target.value;
                      }}
                      value={UserRegister.user.last_name}
                    ></input>
                  </div>
                </div>

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
                        UserRegister.user.email = e.target.value;
                      }}
                      value={UserRegister.user.email}
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
                        UserRegister.user.password = e.target.value;
                      }}
                      value={UserRegister.user.password}
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
                      Create Account
                    </button>
                  </p>

                  <p class="control">
                    <button
                      type="button"
                      class="button is-light"
                      onclick={() => {
                        UserRegister.clear();
                      }}
                    >
                      Clear
                    </button>
                  </p>
                </div>
              </form>
            </div>
          </section>
        </div>
      </main>
    );
  },
};

export default Page;
