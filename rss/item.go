package rss

import (
	"github.com/gr4y/poo/rss/extension"
)

type Item struct {
	extension.ITunesItem
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	PubDate     string    `xml:"pubDate"`
	Link        string    `xml:"link"`
	Guid        string    `xml:"guid"`
	Enclosure   Enclosure `xml:"enclosure"`
}
