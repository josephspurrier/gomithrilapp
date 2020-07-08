import m from "mithril"; // eslint-disable-line no-unused-vars

// Create a flash message class with Bulma.
// http://bulma.io/documentation/components/message/

var View = {
  list: [],
  timeout: 4000, // milliseconds
  prepend: false,
  success: (message) => {
    View.addFlash(message, "is-success");
  },
  failed: (message) => {
    View.addFlash(message, "is-danger");
  },
  warning: (message) => {
    View.addFlash(message, "is-warning");
  },
  primary: (message) => {
    View.addFlash(message, "is-primary");
  },
  link: (message) => {
    View.addFlash(message, "is-link");
  },
  info: (message) => {
    View.addFlash(message, "is-info");
  },
  dark: (message) => {
    View.addFlash(message, "is-dark");
  },
  addFlash: (message, style) => {
    // Don't show a message if zero.
    if (View.timeout === 0) {
      return;
    }

    const msg = {
      message: message,
      style: style,
    };

    //Check if the messages should stack in reverse order.
    if (View.prepend === true) {
      View.list.unshift(msg);
    } else {
      View.list.push(msg);
    }

    m.redraw();

    // Show forever if -1.
    if (View.timeout > 0) {
      setTimeout(() => {
        View.removeFlash(msg);
        m.redraw();
      }, View.timeout);
    }
  },
  removeFlash: (i) => {
    View.list = View.list.filter((v) => {
      return v !== i;
    });
  },
  clear: () => {
    View.list = [];
  },
  view: () => (
    <div style="position: fixed; bottom: 1.5rem; right: 1.5rem; z-index: 100; margin: 0;">
      {View.list.map((i) => (
        <div key={i} class={`notification ${i.style}`}>
          {i.message}
          <button
            class="delete"
            onclick={() => {
              View.removeFlash(i);
            }}
          ></button>
        </div>
      ))}
    </div>
  ),
};

export default View;
