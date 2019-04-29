// Apply this middleware when you only want to allow the page to be accessible
// when the user is NOT authenticated. For instance, the login page.
export default function({ store, redirect }) {
  if (store.getters.isAuthenticated === true) {
    return redirect('/')
  }
}
