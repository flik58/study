import re
from playwright.sync_api import Page, expect

def test_homepage_has_Playwright_in_title_and_get_started_link_linking_to_the_intro_page(page: Page):
    page.goto("https://playwright.dev/")

    expect(page).to_have_title(re.compile("Playwright"))

    get_started = page.locator("text=Get Started")

    expect(get_started).to_have_attribute("href", "/docs/intro")

    get_started.click()

    expect(page).to_have_url(re.compile(".*intro"))
