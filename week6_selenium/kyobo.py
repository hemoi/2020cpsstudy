from selenium.webdriver.common.keys import Keys
from selenium import webdriver
from bs4 import BeautifulSoup
import os
import time

path = os.getcwd() + "/home/hemoi/crawling/chromedriver"
driver = webdriver.Chrome(path)

try :
    # 먼저 검색 창을 찾는다.
    driver.get("http://www.")
    time.sleep()

    element = driver.find_element_by_class_name("main_input")

    # 현재 검색창을 가리킨 상태
    # 이젠 입력을 할 것이다.
    searchIndex = "파이썬"
    element.send_keys(searchIndex)

    # 검색버튼을 찾아야 한다.
    driver.find_element_by_class_name("btn_search").click()

    html = driver.page_source
    bs = BeautifulSoup(html, "html.parser")

    conts = bs.find("div", class_ = "list_search_result").find_all("td", class_ = "detail")

    for c in conts :
        c.find("div", class_ = "title").find("stong").text


finally :
    driver.quit()