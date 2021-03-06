import m from "mithril"; // eslint-disable-line no-unused-vars
import UserRegister from "@/store/userregister";

var Page = ({ attrs }) => {
  // Prefill the fields.
  if (attrs.firstName) {
    UserRegister.user.first_name = attrs.firstName;
  }
  if (attrs.lastName) {
    UserRegister.user.last_name = attrs.lastName;
  }
  if (attrs.email) {
    UserRegister.user.email = attrs.email;
  }
  if (attrs.password) {
    UserRegister.user.password = attrs.password;
  }

  return {
    view: () => {
      return (
        <main>
          <div>
            <section class="section">
              <div class="container">
                <h1 class="title">Register</h1>
                <h2 class="subtitle">Enter your information below.</h2>
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
};

export default Page;
