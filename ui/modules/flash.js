// Create a flash message class with Bulma.
// http://bulma.io/documentation/components/message/
export default class Flash {
  constructor(timeout = 4000, prepend = false, identifier = 'flash-container') {
    this.timeout = timeout
    this.prepend = prepend
    this.identifier = identifier
  }

  success(text) {
    this.showMessage(text, 'is-success')
  }

  failed(text) {
    this.showMessage(text, 'is-danger')
  }

  warning(text) {
    this.showMessage(text, 'is-warning')
  }

  primary(text) {
    this.showMessage(text, 'is-primary')
  }

  link(text) {
    this.showMessage(text, 'is-link')
  }

  info(text) {
    this.showMessage(text, 'is-info')
  }

  dark(text) {
    this.showMessage(text, 'is-dark')
  }

  // showMessage will show the flash message.
  showMessage(text, style) {
    let container = document.getElementById(this.identifier)

    if (container === null) {
      container = document.createElement('div')
      container.id = this.identifier
      container.setAttribute(
        'style',
        `
        position: fixed;
        bottom: 1.5rem;
        right: 1.5rem;
        z-index: 100;
        margin: 0;
        `
      )
      document.body.appendChild(container)
    }

    const el = document.createElement('div')
    el.classList.add('notification', style)

    const btn = document.createElement('button')
    btn.classList.add('delete')
    btn.onclick = () => {
      el.remove()
    }

    el.innerText = text

    el.appendChild(btn)

    // Check if the messages should stack in reverse order.
    if (this.prepend === true) {
      container.insertBefore(el, container.firstChild)
    } else {
      container.appendChild(el)
    }

    // Remove the message after a specific period of time.
    window.setTimeout(() => {
      el.remove()
    }, this.timeout)
  }
}
