package docx

import (
	"bytes"
	"encoding/xml"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"strconv"
)

type Paragraph struct {
	XMLName xml.Name             `xml:"w:p"`
	PPr     *ParagraphProperties `xml:"w:pPr,omitempty"`
	Data    []interface{}

	file *File
}

func (p *Paragraph) AddText(text string) *Run {
	t := &Text{
		Text:     text,
		XMLSpace: "preserve",
	}

	run := &Run{
		Text:          t,
		RunProperties: &RunProperties{},
	}

	p.Data = append(p.Data, run)

	return run
}

func (p *Paragraph) AddLink(text string, link string) *Hyperlink {
	rId := p.file.addLinkRelation(link)
	hyperlink := &Hyperlink{
		ID: rId,
		Run: Run{
			RunProperties: &RunProperties{
				RunStyle: &RunStyle{
					Val: HYPERLINK_STYLE,
				},
			},
			Text: &Text{Text: text, XMLSpace: "preserve"},
		},
	}

	p.Data = append(p.Data, hyperlink)

	return hyperlink
}

func (p *Paragraph) AddImage(data []byte, ext string, widthPx, heightPx int) *Drawing {
	relID, _ := p.file.addImageRelation(ext, data)

	docPrID := p.file.imageIndex
	name := "Image" + strconv.Itoa(docPrID)

	d := &Drawing{
		Inline: &Inline{
			XmlnsWp:  NS_WP,
			XmlnsA:   NS_A,
			XmlnsPic: NS_PIC,
			Extent: &Extent{
				Cx: pxToEMU(widthPx),
				Cy: pxToEMU(heightPx),
			},
			DocPr: &DocPr{
				ID:   docPrID,
				Name: name,
			},
			Graphic: &Graphic{
				GraphicData: &GraphicData{
					Uri: "http://schemas.openxmlformats.org/drawingml/2006/picture",
					Pic: &Pic{
						NvPicPr: &NvPicPr{
							CNvPr: &CNvPr{
								ID:   docPrID,
								Name: name,
							},
							CNvPicPr: &CNvPicPr{},
						},
						BlipFill: &BlipFill{
							Blip:    &Blip{Embed: relID},
							Stretch: &Stretch{FillRect: &FillRect{}},
						},
						SpPr: &SpPr{
							Xfrm: &Xfrm{
								Off: &Off{X: 0, Y: 0},
								Ext: &AExt{
									Cx: pxToEMU(widthPx),
									Cy: pxToEMU(heightPx),
								},
							},
							PrstGeom: &PrstGeom{
								Prst:  "rect",
								AvLst: &AvLst{},
							},
						},
					},
				},
			},
		},
	}

	p.Data = append(p.Data, d)
	return d
}

func (p *Paragraph) AddBreak() {
	type runWithBreak struct {
		XMLName xml.Name       `xml:"w:r"`
		RPr     *RunProperties `xml:"w:rPr,omitempty"`
		Br      *Break         `xml:"w:br"`
	}
	p.Data = append(p.Data, &runWithBreak{Br: &Break{}})
}

func (p *Paragraph) AddImageWithCaption(data []byte, ext string, widthPx, heightPx int, caption string) {
	p.AlignCenter()
	p.AddImage(data, ext, widthPx, heightPx)
	p.AddBreak()
	p.AddText(caption)
}

func (p *Paragraph) AddMultilineText(text string) {
	if text == "" {
		return
	}
	runes := []rune(text)
	start := 0
	for i, ch := range runes {
		if ch == '\n' || ch == '\r' {
			if i > start {
				p.AddText(string(runes[start:i]))
			}
			p.AddBreak()
			start = i + 1
		}
	}
	if start < len(runes) {
		p.AddText(string(runes[start:]))
	}
}

func (p *Paragraph) AddLabelValue(label, value string) {
	if label != "" {
		p.AddText(label + ": ").Bold()
	}
	if value != "" {
		p.AddText(value)
	}
}

func (p *Paragraph) ensurePPr() {
	if p.PPr == nil {
		p.PPr = &ParagraphProperties{}
	}
	if p.PPr.Justification == nil {
		p.PPr.Justification = &Justification{}
	}
}

func (p *Paragraph) AlignLeft() *Paragraph {
	p.ensurePPr()
	p.PPr.Justification.Val = "left"
	return p
}

func (p *Paragraph) AlignCenter() *Paragraph {
	p.ensurePPr()
	p.PPr.Justification.Val = "center"
	return p
}

func (p *Paragraph) AlignRight() *Paragraph {
	p.ensurePPr()
	p.PPr.Justification.Val = "right"
	return p
}

func (p *Paragraph) Justify() *Paragraph {
	p.ensurePPr()
	p.PPr.Justification.Val = "both"
	return p
}

func (p *Paragraph) AddPageBreak() {
	type runWithBreak struct {
		XMLName xml.Name       `xml:"w:r"`
		RPr     *RunProperties `xml:"w:rPr,omitempty"`
		Br      *Break         `xml:"w:br"`
	}
	p.Data = append(p.Data, &runWithBreak{Br: &Break{Type: "page"}})
}

func (p *Paragraph) AddImageAuto(data []byte, ext string, maxWidthPx int) *Drawing {
	cfg, _, err := image.DecodeConfig(bytes.NewReader(data))
	if err != nil || cfg.Width <= 0 || cfg.Height <= 0 {
		return p.AddImage(data, ext, maxWidthPx, int(float64(maxWidthPx)*0.6))
	}
	w := cfg.Width
	h := cfg.Height
	if w > maxWidthPx {
		scale := float64(maxWidthPx) / float64(w)
		w = maxWidthPx
		h = int(float64(h) * scale)
	}
	return p.AddImage(data, ext, w, h)
}
