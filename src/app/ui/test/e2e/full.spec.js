const puppeteer = require('puppeteer')
const timeout = 5000

describe('Index page', () => {
  let browser
  let page

  beforeAll(async () => {
    browser = await puppeteer.launch({
      headless: false
    })
    page = (await browser.pages())[0]
  })

  afterAll(async () => {
    await page.close()
    await browser.close()
  })

  it('registers a user', async () => {
    await page.goto('http://localhost:8080/register', timeout)
    await page.waitForSelector('[name="register"]')
    await page.type('[name="first_name"]', 'John')
    await page.type('[name="last_name"]', 'Doe')
    await page.type('[name="email"]', 'a@a.com')
    await page.type('[name="password"]', 'a')
    await page.click('#submit')
    await page.waitForNavigation()
    const title = await page.title()
    expect(title).toBe('Login')
  })

  it('completes login', async () => {
    await page.goto('http://localhost:8080/login', timeout)
    await page.waitForSelector('[name="login"]')
    await page.type('[name="email"]', 'a@a.com')
    await page.type('[name="password"]', 'a')
    await page.click('#submit')
    await page.waitForNavigation()
    const title = await page.title()
    expect(title).toBe('Welcome')
  })
})
