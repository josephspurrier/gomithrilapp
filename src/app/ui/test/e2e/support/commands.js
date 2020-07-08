/* eslint-disable no-undef */
// ***********************************************
// This example commands.js shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add("login", (email, password) => { ... })
//
//
// -- This is a child command --
// Cypress.Commands.add("drag", { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add("dismiss", { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite("visit", (originalFn, url, options) => { ... })

import Path from "path";

// YOU NEED TO FIX THE WAY THE MIGRATIONS ARE RUN - THE DOCKER WILL NOT EXIST SO RESET WON'T WORK.
// IF IT'S on mac, it should load the stuff and it would work, if it's in travis, the DB will always be clea

Cypress.Commands.add("resetDB", () => {
  return cy
    .exec("MYSQL_ROOT_PASSWORD=password bash ./bash/reset-db.sh", {
      env: { CYPRESS: Path.resolve(__dirname) },
    })
    .its("code")
    .should("eq", 0);
});
