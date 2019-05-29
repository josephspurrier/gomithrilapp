import { mount } from '@vue/test-utils'
import About from '../pages/about.vue'

describe('about.vue', () => {
  it('is setup correctly', () => {
    expect(true).toBe(true)
  })

  it('renders the header', () => {
    const wrapper = mount(About)
    expect(wrapper.text()).toContain('About')
  })
})
