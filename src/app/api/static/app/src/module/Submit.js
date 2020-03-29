var Submit = {
  disabled: false,
  submitText: "Submitting...",
  start: (event) => {
    event.preventDefault();
    this.disabled = true;
  },
  finish: () => {
    this.disabled = false;
  },
  text: (s) => {
    if (!this.disabled) {
      return s;
    } else {
      return this.submitText;
    }
  },
};

export default Submit;
