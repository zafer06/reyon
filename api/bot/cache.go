package bot

import (
	"log"
	"os"
	"io"
	"encoding/json"
	"time"
)

type Cache struct {
	cacheFileName string
}

func (c Cache) CacheData() []Thread {
	c.cacheFileName = "bot/cache.json"

	c.addCache()
	
	var threadList = c.getCache()
	return threadList
}

func (c Cache) addCache() {
	info, err := os.Stat(c.cacheFileName)
	if err != nil {
		var t Thread
		var threadList = t.ParseForum()
		t.SaveData(threadList)
	} else {
		var fileDate = info.ModTime()
		var todayDate = time.Now()

		var diff = todayDate.Sub(fileDate)

		if diff.Minutes() > 60 {
			e := os.Remove(c.cacheFileName)
    		if e != nil {
        		log.Fatal(e)
    		}

			var t Thread
			var threadList = t.ParseForum()
			t.SaveData(threadList)
		}
	}
}

func (c Cache) getCache() []Thread {
	fileContent, err := os.Open(c.cacheFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fileContent.Close()
	
	byteResult, _ := io.ReadAll(fileContent)

	var threadList []Thread
	json.Unmarshal([]byte(byteResult), &threadList)

	return threadList
}