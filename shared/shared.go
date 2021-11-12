package shared

import (
	"encoding/xml"
	"time"
)

type Date time.Time

func (input Date) MarshalXML(element *xml.Encoder, start xml.StartElement) error {
	timeValue := time.Time(input)
	timeString := timeValue.Format("2006-01-02")
	return element.EncodeElement(timeString, start)
}

func (input *Date) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	value, err := ParseDateTime("2006-01-02", decoder, start)

	if err != nil {
		return err
	}

	*input = Date(value)

	return nil
}

func ParseDateTime(
	dateTimeLayout string, decoder *xml.Decoder, start xml.StartElement,
) (time.Time, error) {
	var value string
	err := decoder.DecodeElement(&value, &start)

	if err != nil {
		return time.Time{}, err
	}

	// Use zero time if empty and if MSS
	// returns "0000-00-00"
	if value == "" || value == "0000-00-00" {
		return time.Time{}, nil
	}

	parsed, err := time.Parse(dateTimeLayout, value)

	if err != nil {
		return time.Time{}, err
	}

	return parsed, nil
}

func (date Date) String() string {
	return time.Time(date).String()
}

type LTSData struct {
	A0Ene int `xml:"A0Ene"`
	A0MTV int `xml:"A0MTV"`
	A0Rep int `xml:"A0Rep"`
}

type RoomType int

type HotelType int

type HotelFeature int

type Theme int

type Board int

type OfferType int

type SpecialType int

type SpecialPremium int

type ThemeID int
