package poo

import (
	"github.com/gr4y/poo/lib/extension"
)

type Channel struct {
	extension.ITunesChannel
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
	Language    string `xml:"language"`
	Copyright   string `xml:"copyright"`
	Items       []Item `xml:"item"`
}
