### week4_webCrawling
---
- go_crawler.go
  - 산업보안학과 페이지를 크롤링
    - 제목 [ 학과 ] 를 추출해서 따로 학과라는 태그를 추가함
    - 출력을 제목만 하고 link는 저장
    - 제목에 "Soulick"가 들어간 link만 다시 크롤링해서 본문 내용을 출력
  - _todo_
    - [x] 가끔씩 값이 중복해서 나오는 현상 있음
      - 글의 중복이었음
    - [ ] 뒷부분이 짤리는 부분이 있음

- indeed_crawler.py
  - indeed 웹사이트 크롤링
    - 실습시간에 진행한 beautifulSoup4 각 함수별 기능 정리
    - `file = open("./src/indeed.csv", "a",-1, newline="", encoding='UTF8')` 부분 수정


- indeed.csv
  - indeed_crawler에서 생성된 파일
  
- hyperledger_crawler.py
  - hyperledger 2.0 doc 크롤링
    - 메뉴에 있는 각 링크를 따와서 각 페이지를 크롤링하는 코드작성
    - 실행은 문제없이 되지만 js기반이라 동일한 text만 나옴

- hyperledger_crawling.txt
  - hyperledger_crawler.py 에서 생성된 파일