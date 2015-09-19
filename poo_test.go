package poo_test

import (
	"fmt"
	"github.com/gr4y/poo"
	"io/ioutil"
	"testing"
)

func TestWithURL(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	feed, err := poo.ParseFeed("http://blog.swessel.me/feed")
	if err != nil {
		t.Error(err)
	}
	if feed.Channel.Title != "gr4ys blog" {
		t.Log(feed.Channel.Title)
		t.Fail()
	}
}

func TestWithString(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	xmlFeed := readXML("example-itunes_feed.xml")
	feed, err := poo.ParseXML(xmlFeed)
	if err != nil {
		t.Error(err)
	}
	if feed.Channel.Title != "Back to Work" {
		t.Log(feed.Channel.Title)
		t.Fail()
	}
}

func readXML(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}
