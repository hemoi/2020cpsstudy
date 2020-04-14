import requests
from bs4 import BeautifulSoup
import csv
import sys
class Scraper():
    def __init__(self):
        # url hyperledger fabric 2.0
        self.url = "https://hyperledger-fabric.readthedocs.io/en/release-2.0/"
        self.menuLists = []
        self.rootSoup = self.getHTML(self.url)

        self.menuLists = self.getUrlLists(self.rootSoup)

    def getHTML(self, url):
        # request web page
        res = requests.get(url)

        # request error
        if res.status_code != 200:
            print("request error : ", res.status_code)

        # get html
        html = res.text

        # get information from html
        soup = BeautifulSoup(html, "html.parser")

        return soup
    
    # getUrlLists
    def getUrlLists(self, soup):
        lists = soup.find_all("a", class_ = "reference internal")

        urlLists = []
        for i in lists:
            # if find attrs 'href' then add to lists
            urlLists.append(i.attrs['href'])

        return urlLists

    def getText(self, soup) :
        # title, text

        # mainSoups = soup.find_all("div", class_ = "rst-content").get_text
        mainTexts = soup.find("div", class_ = "rst-content").get_text()

        # for soup in mainSoups:
        #     mainTexts.append(soup.find("p")).text
        # title = mainTexts.select("h1").text
        # parags = texts.find_all("p")

        # want = []

        # for parag in parags:
        #     want.append(parag.text)
        return mainTexts

        

    # main crawling function
    def crawling(self):
        file = open('./hyperledger_crawling.txt', 'w')

        for i in self.menuLists:
            # make full url
            url = self.url + i
            
            file.write(url + "\n")

            # get soup by using full url
            tmpSoup = self.getHTML(url)

            # get Text
            text = self.getText(tmpSoup)

            # write files
            file.write(str(text) + "\n")
        
        file.close()

            




if __name__ == "__main__" :
    s = Scraper()
    s.crawling()
