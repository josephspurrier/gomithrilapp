// Create a flash message class with Bulma.
// http://bulma.io/documentation/components/message/

var Flash = {
  timeout: 4000,
  identifier: "flash-container",
  prepend: false,

  success(text) {
    Flash.showMessage(text, "is-success");
  },

  failed(text) {
    Flash.showMessage(text, "is-danger");
  },

  warning(text) {
    Flash.showMessage(text, "is-warning");
  },

  primary(text) {
    Flash.showMessage(text, "is-primary");
  },

  link(text) {
    Flash.showMessage(text, "is-link");
  },

  info(text) {
    Flash.showMessage(text, "is-info");
  },

  dark(text) {
    Flash.showMessage(text, "is-dark");
  },

  // showMessage will show the flash message.
  showMessage(text, style) {
    // Don't show a message if zero.
    if (Flash.timeout === 0) {
      return;
    }

    let container = document.getElementById(Flash.identifier);
    if (!container) {
      console.log("Could not find flash container.");
      return;
    }

    const el = document.createElement("div");
    el.classList.add("notification", style);

    const btn = document.createElement("button");
    btn.classList.add("delete");
    btn.onclick = () => {
      el.remove();
    };

    el.innerText = text;

    el.appendChild(btn);

    // Check if the messages should stack in reverse order.
    if (Flash.prepend === true) {
      container.insertBefore(el, container.firstChild);
    } else {
      container.appendChild(el);
    }

    // Show forever if -1.
    if (Flash.timeout > 0) {
      window.setTimeout(() => {
        el.remove();
      }, Flash.timeout);
    }
  },
};

export default Flash;
