package poo

import (
	"encoding/xml"
	"fmt"
	"time"
)

type DateTime struct {
	time.Time
}

func (dt *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Define the variables
	var v string
	var err error
	var parsed_time time.Time
	// Decode the Element
	d.DecodeElement(&v, &start)

	// First try RFC1123
	parsed_time, err = time.Parse(time.RFC1123, v)
	if err != nil {
		// If that fails, then try RFC822
		parsed_time, err = time.Parse(time.RFC1123Z, v)
	}
	*dt = DateTime{parsed_time.UTC()}
	return err
}
