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
  text: (s) => (!Submit.disabled ? s : Submit.submitText),
};

export default Submit;
