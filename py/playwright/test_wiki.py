from playwright.sync_api import Playwright, sync_playwright, expect


def run(playwright: Playwright) -> None:
    browser = playwright.chromium.launch(headless=False)
    context = browser.new_context()

    # Open new page
    page = context.new_page()

    # Go to https://www.wikipedia.org/
    page.goto("https://www.wikipedia.org/")

    # Click strong:has-text("日本語")
    page.locator("strong:has-text(\"日本語\")").click()
    page.wait_for_url("https://ja.wikipedia.org/wiki/%E3%83%A1%E3%82%A4%E3%83%B3%E3%83%9A%E3%83%BC%E3%82%B8")

    # Click text=名鉄3400系電車
    page.locator("text=名鉄3400系電車").click()
    page.wait_for_url("https://ja.wikipedia.org/wiki/%E5%90%8D%E9%89%843400%E7%B3%BB%E9%9B%BB%E8%BB%8A")

    # Close page
    page.close()

    # ---------------------
    context.close()
    browser.close()


with sync_playwright() as playwright:
    run(playwright)
