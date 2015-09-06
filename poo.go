package poo

import (
	"encoding/xml"
	"github.com/gr4y/poo/lib"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	feed, err := ParseFeed("http://feeds.5by5.tv/b2w")
	// feed, err := ParseFeed("http://blog.swessel.me/feed")
	// feed, err := ParseFeed("http://blog.swessel.me/feed")
	if err == nil {
		log.Println(feed)
	} else {
		panic(err)
	}
}

func ParseFeed(url string) (poo.Feed, error) {
	var err error
	if isValidURL(url) {
		resp, err := http.Get(url)
		if err == nil {
			b, err := ioutil.ReadAll(resp.Body)
			if err == nil {
				xmlStr := string(b[:])
				return ParseXML(xmlStr)
			}
		}
		return poo.Feed{}, err
	}
	return poo.Feed{}, err
}

func ParseXML(xmlStr string) (poo.Feed, error) {
	feed := poo.Feed{}
	if err := xml.Unmarshal([]byte(xmlStr), &feed); err != nil {
		return feed, err
	}
	return feed, nil
}

func isValidURL(url string) bool {
	matched, err := regexp.MatchString("^http://\\S*", url)
	if err != nil {
		return false
		panic(err)
	}
	return matched
}
