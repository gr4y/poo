package poo_test

import (
	"fmt"
	"github.com/gr4y/poo"
	"io/ioutil"
	"testing"
)

func TestPoo(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	xmlFeed := readXML("_feeds/b2w.xml")
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
