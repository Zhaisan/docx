package docx

import "encoding/xml"

type SectPr struct {
	XMLName xml.Name `xml:"w:sectPr"`
	PgSz    *PgSz    `xml:"w:pgSz,omitempty"`
	PgMar   *PgMar   `xml:"w:pgMar,omitempty"`
}

type PgSz struct {
	W      int    `xml:"w:w,attr"`
	H      int    `xml:"w:h,attr"`
	Orient string `xml:"w:orient,attr,omitempty"`
}
type PgMar struct {
	Top    int `xml:"w:top,attr"`
	Right  int `xml:"w:right,attr"`
	Bottom int `xml:"w:bottom,attr"`
	Left   int `xml:"w:left,attr"`
}

func (f *File) ensureSectPr() *SectPr {
	if f.Document.Body.SectPr == nil {
		f.Document.Body.SectPr = &SectPr{}
	}
	return f.Document.Body.SectPr
}

func mmToTwips(mm int) int {
	return int(float64(mm) * 56.6929133858)
}

func (f *File) SetPageSizeA4() {
	s := f.ensureSectPr()
	s.PgSz = &PgSz{
		W: mmToTwips(210),
		H: mmToTwips(297),
	}
}

func (f *File) SetMarginsMM(top, right, bottom, left int) {
	s := f.ensureSectPr()
	s.PgMar = &PgMar{
		Top: mmToTwips(top), Right: mmToTwips(right),
		Bottom: mmToTwips(bottom), Left: mmToTwips(left),
	}
}
