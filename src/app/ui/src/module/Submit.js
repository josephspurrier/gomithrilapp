var Submit = {
  disabled: false,
  submitText: "Submitting...",
  start: (event) => {
    event.preventDefault();
    Submit.disabled = true;
  },
  finish: () => {
    Submit.disabled = false;
  },
  text: (s) => {
    if (!Submit.disabled) {
      return s;
    } else {
      return Submit.submitText;
    }
  },
};

export default Submit;
