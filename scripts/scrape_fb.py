# bazel run //scripts:scrape_fb
import asyncio
import os
from typing import List, cast

from dotenv import find_dotenv, load_dotenv
from pyppeteer import launch
from pyppeteer.element_handle import ElementHandle

load_dotenv()
print('.env loaded from:', find_dotenv())

async def get_text(root: ElementHandle, xpath: str, property: str) -> str:
  nodes = await root.xpath(xpath)
  assert nodes, "no xpath found for " + xpath
  val = await nodes[0].getProperty(property)
  return cast(str, await val.jsonValue())

async def get_text_selector(root: ElementHandle, selector: str, property: str) -> str:
  node = await root.querySelector(selector)
  assert node, "no selector found for " + selector 
  val = await node.getProperty(property)
  return cast(str, await val.jsonValue())

async def get_texts(root: ElementHandle, xpath: str, property: str) -> List[str]:
  nodes = await root.xpath(xpath)
  vals = await asyncio.gather(*[node.getProperty(property) for node in nodes])
  return [cast(str, await val.jsonValue()) for val in vals]

async def main():
    browser = await launch(headless=False, autoClose=False)
    await browser._connection.send('Browser.grantPermissions', {
      'origin': "https://www.facebook.com",
      'permissions': ['notifications'],
    })
    page = await browser.newPage()
    await page.goto('https://www.facebook.com/')
    await page.click('#email')
    await page.keyboard.type(os.environ['FB_USERNAME'])
    await page.click('#pass')
    await page.keyboard.type(os.environ['FB_PASSWORD'])
    await asyncio.wait([
        asyncio.create_task(page.keyboard.press('Enter')),
        asyncio.create_task(page.waitForNavigation()),
    ])
    await asyncio.wait([
      asyncio.create_task(page.goto('https://www.facebook.com/groups/1028477820535630')),
      asyncio.create_task(page.waitForNavigation()),
    ])
    await page.waitFor(2000)
    await page.evaluate("{window.scrollBy(0, document.body.scrollHeight);}")
    await page.waitForXPath('//div[@role="article"]//div[@data-ad-preview="message"]/../..//div[text()="See More"]')
    posts = await page.xpath('//div[@role="article"]//div[@data-ad-preview="message"]/../..')
    while see_more := await page.xpath('//div[text()="See More"]'):
      # for whatever reason, it can't click "See More" in the group "About" section
      if len(see_more) == 1: break
      print('clicking on ', see_more[0])
      await see_more[0].click(delay=300)
      await page.waitFor(1000)
    for post in posts:
      link = get_text(post, '//a[contains(@href,"/commerce/listing")]', 'href')
      pics = get_texts(post, '//a[contains(@href,"/commerce/listing")]//img', 'src')
      main_text = get_text_selector(post, ':first-child', 'textContent')
      bottom_text = get_text_selector(post, ':last-child', 'innerText')
      link, pics, main_text, bottom_text = await asyncio.gather(link, pics, main_text, bottom_text)
      info = {
        "link": link, 
        "pics": pics, 
        "main_text": main_text, 
        "bottom_text": bottom_text
      }
      print(info)

asyncio.get_event_loop().run_until_complete(main())
