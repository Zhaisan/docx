package docx

import "encoding/xml"

const (
	HYPERLINK_STYLE = "a1"
)

type RunProperties struct {
	XMLName  xml.Name  `xml:"w:rPr"`
	Color    *Color    `xml:"w:color,omitempty"`
	Size     *Size     `xml:"w:sz,omitempty"`
	RunStyle *RunStyle `xml:"w:rStyle,omitempty"`

	RFonts    *RunFonts  `xml:"w:rFonts,omitempty"`
	Bold      *Bold      `xml:"w:b,omitempty"`
	Italic    *Italic    `xml:"w:i,omitempty"`
	Underline *Underline `xml:"w:u,omitempty"`
}

type RunStyle struct {
	XMLName xml.Name `xml:"w:rStyle"`
	Val     string   `xml:"w:val,attr"`
}

type Color struct {
	XMLName xml.Name `xml:"w:color"`
	Val     string   `xml:"w:val,attr"`
}

type Size struct {
	XMLName xml.Name `xml:"w:sz"`
	Val     int      `xml:"w:val,attr"`
}

type RunFonts struct {
	XMLName  xml.Name `xml:"w:rFonts"`
	Ascii    string   `xml:"w:ascii,attr,omitempty"`
	HAnsi    string   `xml:"w:hAnsi,attr,omitempty"`
	CS       string   `xml:"w:cs,attr,omitempty"`
	EastAsia string   `xml:"w:eastAsia,attr,omitempty"`
}

type Bold struct {
	XMLName xml.Name `xml:"w:b"`
}
type Italic struct {
	XMLName xml.Name `xml:"w:i"`
}

type Underline struct {
	XMLName xml.Name `xml:"w:u"`
	Val     string   `xml:"w:val,attr,omitempty"`
}
