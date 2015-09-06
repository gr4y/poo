package rss

type Enclosure struct {
	URL      string `xml:"url,attr"`
	Length   string `xml:"length,attr"`
	MimeType string `xml:"type,attr"`
}
