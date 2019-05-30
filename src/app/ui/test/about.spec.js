import { mount } from '@vue/test-utils'
import About from '../pages/about.vue'

describe('about.vue', () => {
  const wrapper = mount(About)

  it('is setup correctly', () => {
    expect(true).toBe(true)
  })

  it('renders text', () => {
    expect(wrapper.text()).toContain('About')
    expect(wrapper.text()).toContain('This shows you how to build a website')
  })

  it('has a few headings', () => {
    expect(wrapper.contains('h1')).toBe(true)
    expect(wrapper.contains('h2')).toBe(true)
  })
})
