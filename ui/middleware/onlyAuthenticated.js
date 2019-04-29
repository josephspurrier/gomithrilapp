// Apply this middleware when you only want to allow the page to be accessible
// when the user is authenticated. For instance, a private page.
export default function({ route, store, redirect }) {
  // Whitelist routes that don't require authentication.
  for (const url of ['/login', '/about']) {
    if (route.path === url) {
      return
    }
  }

  // Redirect the to login page if not authenticated.
  if (store.getters.isAuthenticated !== true) {
    return redirect('/login')
  }
}
