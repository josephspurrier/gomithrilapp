module.exports = {
  disabled: false,
  submitText: 'Submitting...',
  start: function(event) {
    event.preventDefault()
    this.disabled = true
  },
  finish: function() {
    this.disabled = false
  },
  text: function(s) {
    if (!this.disabled) {
      return s
    } else {
      return this.submitText
    }
  }
}