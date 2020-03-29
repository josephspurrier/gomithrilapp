import Cookie from "js-cookie";

var Auth = {
  cookieName: "auth",
  save: (auth) => {
    Cookie.set(Auth.cookieName, auth);
  },
  clear: () => {
    Cookie.remove(Auth.cookieName);
  },
  bearerToken: () => {
    let auth = Cookie.get("auth");
    if (auth === undefined) {
      return false;
    }

    let v = JSON.parse(auth);
    return "Bearer " + v.accessToken;
  },
  isLoggedIn: () => {
    try {
      let auth = Cookie.get(Auth.cookieName);
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

export default Auth;
