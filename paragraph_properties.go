package docx

import "encoding/xml"

type ParagraphProperties struct {
	XMLName       xml.Name       `xml:"w:pPr"`
	Justification *Justification `xml:"w:jc,omitempty"`
}

type Justification struct {
	XMLName xml.Name `xml:"w:jc"`
	Val     string   `xml:"w:val,attr"`
}
