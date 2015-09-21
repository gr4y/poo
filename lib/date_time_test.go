package poo_test

import (
	"fmt"
	"github.com/gr4y/poo"
	"io/ioutil"
	"testing"
)

func TestRFC1132(t *testing.T) {
	// if testing.Short() {
	// 	t.Skip("Skipping test in short mode.")
	// }
	xmlFeed := readXML("../feeds/example-rfc1132.xml")
	feed, err := poo.ParseXML(xmlFeed)
	if len(feed.Channel.Items) < 3 {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
}

func TestRFC1132Z(t *testing.T) {
	// if testing.Short() {
	// 	t.Skip("Skipping test in short mode.")
	// }
	xmlFeed := readXML("../feeds/example-rfc1132z.xml")
	feed, err := poo.ParseXML(xmlFeed)
	if len(feed.Channel.Items) < 3 {
		t.Fail()
	}

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
