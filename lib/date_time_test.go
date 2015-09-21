package poo_test

import (
	"fmt"
	"github.com/gr4y/poo"
	"io/ioutil"
	"testing"
)

func TestRFC1123(t *testing.T) {
	// if testing.Short() {
	// 	t.Skip("Skipping test in short mode.")
	// }
	xmlFeed := readXML("example/rfc1123-feed.xml")
	feed, err := poo.ParseXML(xmlFeed)
	t.Log(feed.Channel.Items[0].PubDate)
	if err != nil {
		t.Error(err)
	}
}

func TestRFC1123Z(t *testing.T) {
	// if testing.Short() {
	// 	t.Skip("Skipping test in short mode.")
	// }
	xmlFeed := readXML("example/rfc1123z-feed.xml")
	feed, err := poo.ParseXML(xmlFeed)
	t.Log(feed.Channel.Items[0].PubDate)
	if err != nil {
		t.Error(err)
	}
}

func readXML(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}
