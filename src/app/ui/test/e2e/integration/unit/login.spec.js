import { createLocalVue } from '@vue/test-utils'
import VeeValidate from 'vee-validate'
import Login from '@/pages/login.vue'

const mount = require('cypress-vue-unit-test')

describe('login.vue', function() {
  // Create an extended `Vue` constructor.
  const localVue = createLocalVue()
  localVue.use(VeeValidate)
  beforeEach(mount(Login /*, {
      localVue: localVue
    } */))

  /* const wrapper = mount(Login, {
    localVue: localVue
  }) */

  it('renders text', function() {
    cy.wrap(Cypress.vue).should('not.be.undefined')
  })

  it('has a few headings', function() {
    cy.get('h1').contains('Login')
    cy.get('h2').contains('Enter your login information below.')
  })

  it('should not allow login when fields are blank', function() {
    cy.get('[data-cy=email]')
      .type('a@a.com')
      .should('have.value', 'a@a.com')

    cy.get('[data-cy=password]')
      .type('a')
      .should('have.value', 'a')

    cy.get('[data-cy=submit]').click()

    cy.contains('Login f.')
  })

  /*

  it('should not allow login when fields are blank', () => {
    expect(wrapper.vm.login.email).toBe('')
    expect(wrapper.vm.login.password).toBe('')
    wrapper.vm.login.email = 'a@a.com'
    wrapper.vm.login.password = 'a'
    const button = wrapper.find('#submit')
    button.trigger('click')
    // debugger
    // expect(wrapper.vm.count).toBe(1)
  }) */
})
