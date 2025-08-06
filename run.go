package docx

import "encoding/xml"

type Run struct {
	XMLName       xml.Name       `xml:"w:r"`
	RunProperties *RunProperties `xml:"w:rPr,omitempty"`
	InstrText     string         `xml:"w:instrText,omitempty"`
	Text          *Text

	Drawing *Drawing `xml:"w:drawing,omitempty"`
}

type Text struct {
	XMLName  xml.Name `xml:"w:t"`
	XMLSpace string   `xml:"xml:space,attr,omitempty"`
	Text     string   `xml:",chardata"`
}

func (r *Run) ensureProps() {
	if r.RunProperties == nil {
		r.RunProperties = &RunProperties{}
	}
}

func (r *Run) Color(color string) *Run {
	r.ensureProps()

	r.RunProperties.Color = &Color{
		Val: color,
	}

	return r
}

func (r *Run) Size(size int) *Run {
	r.ensureProps()

	r.RunProperties.Size = &Size{
		Val: size * 2,
	}
	return r
}

func (r *Run) Font(name string) *Run {
	r.ensureProps()
	r.RunProperties.RFonts = &RunFonts{
		Ascii: name, HAnsi: name, CS: name, EastAsia: name,
	}
	return r
}

func (r *Run) Bold() *Run {
	r.ensureProps()
	r.RunProperties.Bold = &Bold{}
	return r
}

func (r *Run) Italic() *Run {
	r.ensureProps()
	r.RunProperties.Italic = &Italic{}
	return r
}

func (r *Run) Underline() *Run {
	r.ensureProps()
	r.RunProperties.Underline = &Underline{Val: "single"}
	return r
}

type Hyperlink struct {
	XMLName xml.Name `xml:"w:hyperlink"`
	ID      string   `xml:"r:id,attr"`
	Run     Run
}

type Break struct {
	XMLName xml.Name `xml:"w:br"`
	Type    string   `xml:"w:type,attr,omitempty"`
}
