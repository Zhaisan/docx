package docx

import "encoding/xml"

const (
	NS_A   = "http://schemas.openxmlformats.org/drawingml/2006/main"
	NS_WP  = "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
	NS_PIC = "http://schemas.openxmlformats.org/drawingml/2006/picture"
)

const emuPerPixel = 9525

func pxToEMU(px int) int64 { return int64(px) * emuPerPixel }

type Drawing struct {
	XMLName xml.Name `xml:"w:drawing"`
	Inline  *Inline  `xml:"wp:inline"`
}

type Inline struct {
	XMLName  xml.Name `xml:"wp:inline"`
	XmlnsWp  string   `xml:"xmlns:wp,attr,omitempty"`
	Extent   *Extent  `xml:"wp:extent"`
	DocPr    *DocPr   `xml:"wp:docPr"`
	Graphic  *Graphic `xml:"a:graphic"`
	XmlnsA   string   `xml:"xmlns:a,attr,omitempty"`
	XmlnsPic string   `xml:"xmlns:pic,attr,omitempty"`
}

type Extent struct {
	Cx int64 `xml:"cx,attr"`
	Cy int64 `xml:"cy,attr"`
}

type DocPr struct {
	ID   int    `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type Graphic struct {
	XMLName     xml.Name     `xml:"a:graphic"`
	GraphicData *GraphicData `xml:"a:graphicData"`
}

type GraphicData struct {
	XMLName xml.Name `xml:"a:graphicData"`
	Uri     string   `xml:"uri,attr"`
	Pic     *Pic     `xml:"pic:pic"`
}

type Pic struct {
	XMLName  xml.Name  `xml:"pic:pic"`
	NvPicPr  *NvPicPr  `xml:"pic:nvPicPr"`
	BlipFill *BlipFill `xml:"pic:blipFill"`
	SpPr     *SpPr     `xml:"pic:spPr"`
}

type NvPicPr struct {
	CNvPr    *CNvPr    `xml:"pic:cNvPr"`
	CNvPicPr *CNvPicPr `xml:"pic:cNvPicPr"`
}

type CNvPr struct {
	ID   int    `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

type CNvPicPr struct{}

type BlipFill struct {
	Blip    *Blip    `xml:"a:blip"`
	Stretch *Stretch `xml:"a:stretch"`
}

type Blip struct {
	Embed string `xml:"r:embed,attr"`
}

type Stretch struct {
	FillRect *FillRect `xml:"a:fillRect"`
}
type FillRect struct{}

type SpPr struct {
	Xfrm     *Xfrm     `xml:"a:xfrm"`
	PrstGeom *PrstGeom `xml:"a:prstGeom"`
}
type PrstGeom struct {
	Prst  string `xml:"prst,attr"`
	AvLst *AvLst `xml:"a:avLst"`
}
type AvLst struct{}

type Xfrm struct {
	Off *Off  `xml:"a:off"`
	Ext *AExt `xml:"a:ext"`
}
type Off struct {
	X int64 `xml:"x,attr"`
	Y int64 `xml:"y,attr"`
}
type AExt struct {
	Cx int64 `xml:"cx,attr"`
	Cy int64 `xml:"cy,attr"`
}
