// Apply this middleware when you only want to allow the page to be accessible
// when the user is authenticated. For instance, a private page.
export default function({ store, redirect }) {
  if (store.getters.isAuthenticated !== true) {
    return redirect('/login')
  }
}
