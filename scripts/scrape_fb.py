import asyncio
import os

from dotenv import load_dotenv, find_dotenv
from pyppeteer import launch

load_dotenv()
print('.env:', os.environ, find_dotenv())

async def main():
    browser = await launch(headless=False, autoClose=False)
    page = await browser.newPage()
    await page.goto('https://www.facebook.com/')
    await page.click('#email')
    await page.keyboard.type(os.environ['FB_USERNAME'])
    await page.click('#pass')
    await page.keyboard.type(os.environ['FB_PASSWORD'])
    await page.screenshot({'path': 'google.png'})
    await asyncio.wait([
        asyncio.create_task(page.keyboard.press('Enter')),
        asyncio.create_task(page.waitForNavigation()),
    ])
    await asyncio.wait([
      asyncio.create_task(page.goto('https://www.facebook.com/groups/1028477820535630')),
      asyncio.create_task(page.waitForNavigation()),
    ])
    posts = await page.xpath('//div[@role="article"]//div[@data-ad-preview="message"]/../..')
    for post in posts:
      links = await post.xpath('//div[text()="See More"]')
      pics = await post.xpath('//a[contains(@href,"/commerce/listing")]//img')
      if links:
        await links[0].click(delay=576)
      


asyncio.get_event_loop().run_until_complete(main())
