from selenium.webdriver.common.keys import Keys
from selenium import webdriver
from bs4 import BeautifulSoup
import urllib.request
import os
import time

basepath = os.getcwd() + "/week6_selenium"
driver_path = basepath + "/chromedriver"
img_path = basepath + "/img/"

# if os.path.isdir("./img") == False:
#     os.mkdir(basepath+"/img")


driver = webdriver.Chrome(driver_path)
keyword = "ibm"

# to find ImgFomat
def findImgFormat(imgSrc):
    if ".png" in imgSrc:
        return ".png"

    elif ".jpg" in imgSrc:
        return ".jpg"
    
    elif ".jpeg" in imgSrc:
        return ".jpeg"
    
    elif ".gif" in imgSrc:
        return ".gif"

    else:
        return
    
try:
    # 네이버 사전이동
    driver.get("https://terms.naver.com/")
    time.sleep(1)

    # 현재 검색창을 가리킨 상태
    element = driver.find_element_by_id("term_query")

    searchIndex = keyword
    element.send_keys(searchIndex)

    # 검색
    driver.find_element_by_xpath('//*[@id="terms_search_form"]/fieldset/div/input[1]').click()
    time.sleep(3)

    html = driver.page_source
    bs = BeautifulSoup(html, "html.parser")

    img_src = bs.find("div", class_ = "thumb_area").find("img").get("src")
    imgFormat = findImgFormat(img_src)
    urllib.request.urlretrieve(img_src, img_path + searchIndex + imgFormat)

except:
    print("error!")
    driver.quit()

finally:
    print("quit Driver")
    time.sleep(2)
    driver.quit()
