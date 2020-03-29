var m = require('mithril')

var data = {
  title: 'Login',
  subtitle: 'Enter your login information below.'
}

module.exports = {
  view: () =>
<main>
  <div>
    <section class="section">
      <div class="container">
        <h1 class="title">{ data.title }</h1>
        <h2 class="subtitle">{ data.subtitle }</h2>
      </div>

      <div class="container" style="margin-top: 1em;">
          <form name="login">

            <div class="field">
              <label class="label">Email</label>
              <div class="control">
                <input label="Email" name="email" type="text" class="input" data-cy="email" required></input>
              </div>
            </div>

            <div class="field">
              <label class="label">Password</label>
              <div class="control">
                <input label="Password" name="password" type="password" class="input" data-cy="password" required></input>
              </div>
            </div>

            <div class="field is-grouped">
              <p class="control">
                <button id="submit" type="submit" data-cy="submit" class="button is-primary">
                  Submit
                </button>
              </p>

              <p class="control">
                <button type="button" class="button is-light">
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
}