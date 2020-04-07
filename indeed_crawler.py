import requests
from bs4 import BeautifulSoup
import csv

class Scraper():
    def __init__(self):
        self.url = "https://kr.indeed.com/jobs?q=python&limit=50"

    def getHTML(self, cnt):        
        # 웹사이트 접속
        res = requests.get(self.url + "&start=" + str(cnt * 50))

        # 만약 정상적으로 접속하지 않았다면
        if res.status_code != 200 :
            print("request error : ", res.status_code)

        # 웹사이트 html 받아오기
        html = res.text

        # html 에서 원하는 정보 얻기
        soup = BeautifulSoup(html, "html.parser")

        return soup

    def getPages(self, soup) :
        # .은 class를 선택하는 것을 말한다.
        # span.pagination 등등
        # pagination > a는 a태그를 가져 오겠다
        pages = soup.select(".pagination > a")

        # 몇개의 페이지가 있는지 가져오겠다
        return len(pages)

    def getCards(self, soup, cnt):
        jobCards = soup.find_all("div", class_ = "jobsearch-SerpJobCard")

        jobID = []
        jobTitle = []
        jobLocation = []

        for job in jobCards:
            jobTitle.append(job.find("a").text.replace("\n",""))
            # non-type이기 때문에 반환할 수 없다
            # jobLocation.append(job.find("div", class_ = "location").text)
            if job.find("div", class_ = "location") != None :
                jobLocation.append(job.find("div", class_ = "location").text)
            elif job.find("span", class_ = "location") != None: 
                jobLocation.append(job.find("span", class_ = "location").text)
            jobID.append("https://kr.indeed.com/viewjob?jk=" + job["data-jk"])

        self.writeCSV(jobID, jobTitle, jobLocation, cnt)

    def writeCSV(self, jobID, jobTitle, jobLocation, cnt):
        file = open("./src/indeed.csv", "a",-1, newline="", encoding='UTF8')

        wr = csv.writer(file)
        for i in range(len(jobID)) :
            wr.writerow([str(i + 1 +(cnt*50)), jobID[i], jobTitle[i], jobLocation[i]])

        file.close

    def scrap(self):
        soupPage = self.getHTML(0)
        pages = self.getPages(soupPage)

        # 초기화를 시켜준다
        file = open("./src/indeed.csv", "w",-1, newline="", encoding='UTF8')
        wr = csv.writer(file)
        wr.writerow(["No.","Link", "Title", "Location"])
        file.close

        for i in range(pages) :
            soupCard = self.getHTML(i)
            self.getCards(soupCard, i)
            print(i + 1, "번째 페이지 Done")

if __name__ == "__main__" :
    s = Scraper()
    s.scrap()