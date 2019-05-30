const puppeteer = require('puppeteer')
const timeout = 5000

describe('Index page', () => {
  it('is setup correctly', () => {
    expect(true).toBe(true)
  })

  let browser
  let page
  beforeAll(async () => {
    browser = await puppeteer.launch({
      headless: false
    })
    // page = await browser.newPage()
    page = (await browser.pages())[0]
    await page.goto('http://localhost:8080/login', timeout)
  })

  afterAll(async () => {
    await page.close()
    await browser.close()
  })

  it('renders login page', async () => {
    await page.type('[name="email"]', 'a@a.com')
    await page.type('[name="password"]', 'a')
    await page.click('#submit')
    await page.waitForNavigation()
    /* await page.screenshot({
      path: './test/screenshot/1.png'
    }) */
  })
})
