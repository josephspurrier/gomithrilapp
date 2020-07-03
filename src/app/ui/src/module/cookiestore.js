import Cookie from "js-cookie";

var CookieStore = {
  cookieName: "auth",
  save: (auth) => {
    Cookie.set(CookieStore.cookieName, auth);
  },
  clear: () => {
    Cookie.remove(CookieStore.cookieName);
  },
  bearerToken: () => {
    let auth = Cookie.get(CookieStore.cookieName);
    if (auth === undefined) {
      return false;
    }

    let v = JSON.parse(auth);
    return "Bearer " + v.accessToken;
  },
  isLoggedIn: () => {
    try {
      let auth = Cookie.get(CookieStore.cookieName);
      if (auth === undefined) {
        return false;
      }
      return true;
    } catch (err) {
      console.log(err);
    }

    return false;
  },
};

export default CookieStore;
