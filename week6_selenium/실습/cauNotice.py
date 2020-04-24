from selenium import webdriver
from bs4 import BeautifulSoup
import os
import time

path = os.getcwd() + "/week6_selenium"

#  chrome driver가 필요하다
driver = webdriver.Chrome(path)

# try :
#     driver.get("")
#     time.sleep(1)
#     #driver.implicitly_wait(10) : 로딩 끝날 때까지 기다린 후에 (최대 10초) 시작한다!

#     html = driver.page_source
#     bs = BeautifulSoup(html, "html.parser")

#     pages = bs.find("div", class_= "pagination").find_all("a")[-1]["href"].split("page")[1]
#     pages = int(pages)
    
#     title = []
#     for i in range(pages) :
#         driver.get("https://www.cau.ac.kr/cms/FR_CON/index.do?MENU=ID=2130#page" + str(i + 1))
#         time.sleep(3)

#         html = driver.page_source
#         bs = BeautifulSoup(html, "html.parser")

#         conts = bs.find_all("div", class_ = "txtL")
#         for c in conts :
#             print(c.find("a").text)

try :
    driver.get("")
    time.sleep(1)
    #driver.implicitly_wait(10) : 로딩 끝날 때까지 기다린 후에 (최대 10초) 시작한다!

    html = driver.page_source
    bs = BeautifulSoup(html, "html.parser")

    
    title = []
    for i in range(3) :
        driver.get("https://www.cau.ac.kr/cms/FR_CON/index.do?MENU=ID=2130#page" + str(i + 1))
        time.sleep(3)

        html = driver.page_source
        bs = BeautifulSoup(html, "html.parser")

        conts = bs.find_all("div", class_ = "txtL")
        title.append("page" + str(i + 1))
        for c in conts :
            title.append(print(c.find("a").text))


finally:
    for t in title :
        # 값이 존재하면 몇번째에 있는지를 알려준다. 즉 없지 않은 경우에만 한다는 것을 말한다.
        if t.find("page") != -1:
            print()
            print(t)
        else:
            print(t)
    time.sleep(3)
    driver.quit()