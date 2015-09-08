package extension

type ITunesItem struct {
	Author   string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd author"`
	Subtitle string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd subtitle"`
	Summary  string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd summary"`
	Explicit string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd explicit"`
	Duration string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd duration"`
	Image    string `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd image"`
}
