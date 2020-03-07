const puppeteer = require('puppeteer')
const helper = require('@/test/helper.js')

describe('full application interaction', () => {
  const timeout = 5000
  const baseURL = 'http://localhost:8080'
  let browser
  let page

  beforeAll(async () => {
    browser = await puppeteer.launch({
      // slowMo: 50,
      headless: false
    })
    page = (await browser.pages())[0]
    page.setViewport({ width: 1200, height: 800 })
  })

  afterAll(async () => {
    await page.close()
    await browser.close()
  })

  it('registers a user', async () => {
    await page.goto(baseURL + '/register', timeout)
    await page.waitForSelector('[name="register"]')
    await page.type('[name="first_name"]', 'John')
    await page.type('[name="last_name"]', 'Doe')
    await page.type('[name="email"]', 'a@a.com')
    await page.type('[name="password"]', 'a')
    await page.click('#submit')
    // await page.waitForSelector('title')
    await page.waitFor(1000)
    const title = await page.title()
    expect(title).toBe('Login')
  })

  it('completes login', async () => {
    await page.goto(baseURL + '/login', timeout)
    await page.waitForSelector('[name="login"]')
    await page.type('[name="email"]', 'a@a.com')
    await page.type('[name="password"]', 'a')
    await page.click('#submit')
    await page.waitForSelector('title')
    const title = await page.title()
    expect(title).toBe('Welcome')
  })

  it('add a note', async () => {
    // Load the page
    await page.goto(baseURL + '/note', timeout)
    await page.waitForSelector('#note-section')

    // Create an item.
    await page.type('[name="note-add"]', 'This is a note.')
    page.keyboard.press('Enter')

    // Wait for the item to appear.
    await page.waitForSelector('.fa-trash-o')
    const count = (await page.$$('.fa-trash-o')).length
    expect(count).toBe(1)
  })

  it('edit a note', async () => {
    // Load the page
    await page.goto(baseURL + '/note', timeout)
    await page.waitForSelector('#note-section')

    // Edit an item.
    await page.type('.individual-note', 'This is a note edit.')
    page.keyboard.press('Enter')

    // Wait for the item to appear.
    await page.waitForSelector('.fa-trash-o')
    const count = (await page.$$('.fa-trash-o')).length
    expect(count).toBe(1)
  })

  it('delete a note', async () => {
    // Load the page.
    await page.goto(baseURL + '/note', timeout)
    await page.waitForSelector('#note-section')

    // Delete the item.
    await page.click('.fa-trash-o')
    await helper.delay(250)
    const count = (await page.$$('.fa-trash-o')).length
    expect(count).toBe(0)
  })
})
