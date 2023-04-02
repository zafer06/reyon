package bot

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
)

type Thread struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	User        string `json:"user"`
	Date        string `json:"date"`
	Reply       string `json:"reply"`
	Description string `json:"description"`
}

func (t Thread) ParseForum() []Thread {
	var html = t.requestHTML()
	//var html = t.requestFile("reyon-jobs.html")

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		log.Fatal(err)
	}

	var threadList []Thread

	// Find the review items
	doc.Find("#inlinemodform .thread ol").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		div := s.Find("div.title")
		title := div.Find("a").Text()
		link, _ := div.Find("a").Attr("href")
		desc, _ := div.Find("a").Find("span").Attr("data-original-title")

		user := s.Find(".user .desktop a").Text()
		date := s.Find(".date").Text()
		reply := s.Find(".user span a").First().Text()

		t.Title = title
		t.Link = link
		t.User = user
		t.Date = date
		t.Reply = reply
		t.Description = desc

		threadList = append(threadList, t)
	})

	// Verileri cache icin sakla
	t.SaveData(threadList)

	return threadList
}

func (t Thread) SaveData(list []Thread) error {
	threadList, _ := json.Marshal(list)
	return ioutil.WriteFile("bot/cache.json", threadList, 0644)
}

func (t Thread) requestHTML() io.Reader {
	// Request the HTML page.
	res, err := http.Get("https://www.r10.net/kodlama-isi-veren-yazilim-firmalari/")
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	utfBody, _ := iconv.NewReader(res.Body, "windows-1254", "utf-8")
	//newBody, _ := ioutil.ReadAll(utfBody)
	//bodyStr := string(newBody)
	//fmt.Println(bodyStr)

	return utfBody
}

func (t Thread) requestFile(address string) io.Reader {
	file, _ := os.Open(address)
	return file
}

// <script src="https://cdn.r10.net/modern/js/plugins.js?v=7.4.22&amp;rand=2"></script>
