package extension

type ITunesChannel struct {
	Author   string      `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd author"`
	Keywords string      `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd keywords"`
	Explicit string      `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd explicit"`
	Image    ITunesImage `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd image"`
	Block    string      `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd block"`
	Category string      `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd category"`
	Owner    ITunesOwner `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd owner"`
}
