import { mount, createLocalVue } from '@vue/test-utils'
import VeeValidate from 'vee-validate'
import Login from '@/pages/login.vue'

describe('login.vue', () => {
  // Create an extended `Vue` constructor.
  const localVue = createLocalVue()
  localVue.use(VeeValidate)
  const wrapper = mount(Login, {
    localVue: localVue
  })

  it('is setup correctly', () => {
    expect(true).toBe(true)
  })

  it('renders text', () => {
    expect(wrapper.text()).toContain('Login')
    expect(wrapper.text()).toContain('Enter your login information below.')
  })

  it('has a few headings', () => {
    expect(wrapper.contains('h1')).toBe(true)
    expect(wrapper.contains('h2')).toBe(true)
  })

  it('should not allow login when fields are blank', () => {
    expect(wrapper.vm.login.email).toBe('')
    expect(wrapper.vm.login.password).toBe('')
    wrapper.vm.login.email = 'a@a.com'
    wrapper.vm.login.password = 'a'
    const button = wrapper.find('#submit')
    button.trigger('click')
    // debugger
    // expect(wrapper.vm.count).toBe(1)
  })
})
