package docx

import (
	"archive/zip"
	"encoding/xml"
	"io"
	"os"
	"strconv"
)

type mediaFile struct {
	path        string
	contentType string
	data        []byte
}

type File struct {
	Document    Document
	DocRelation DocRelation

	rId        int
	imageIndex int
	media      []mediaFile

	stylesXML string
}

func NewFile() *File {
	defaultRel := []*RelationShip{
		{
			ID:     "rId1",
			Type:   `http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles`,
			Target: "styles.xml",
		},
		{
			ID:     "rId2",
			Type:   `http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme`,
			Target: "theme/theme1.xml",
		},
		{
			ID:     "rId3",
			Type:   `http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable`,
			Target: "fontTable.xml",
		},
	}

	f := &File{
		Document: Document{
			XMLName: xml.Name{
				Space: "w",
			},
			XMLW: XMLNS_W,
			XMLR: XMLNS_R,
			Body: &Body{
				XMLName: xml.Name{
					Space: "w",
				},
				Paragraph: make([]*Paragraph, 0),
			},
		},
		DocRelation: DocRelation{
			Xmlns:        XMLNS,
			Relationship: defaultRel,
		},
		rId:       4,
		stylesXML: TEMP_WORD_STYLE,
	}

	return f
}

func (f *File) Save(path string) (err error) {
	fzip, _ := os.Create(path)
	defer fzip.Close()

	zipWriter := zip.NewWriter(fzip)
	defer zipWriter.Close()

	return f.pack(zipWriter)
}

func (f *File) Write(writer io.Writer) (err error) {
	zipWriter := zip.NewWriter(writer)
	defer zipWriter.Close()

	return f.pack(zipWriter)
}

func (f *File) AddParagraph() *Paragraph {
	p := &Paragraph{
		Data: make([]interface{}, 0),
		file: f,
	}

	f.Document.Body.Paragraph = append(f.Document.Body.Paragraph, p)
	return p
}

func (f *File) addLinkRelation(link string) string {
	rel := &RelationShip{
		ID:         "rId" + strconv.Itoa(f.rId),
		Type:       REL_HYPERLINK,
		Target:     link,
		TargetMode: REL_TARGETMODE,
	}

	f.rId += 1

	f.DocRelation.Relationship = append(f.DocRelation.Relationship, rel)

	return rel.ID
}

func (f *File) UseTimesNewRomanDefault() {
	f.stylesXML = TEMP_WORD_STYLE_TNR
}

func (f *File) pack(zipWriter *zip.Writer) (err error) {
	files := map[string]string{}

	files["_rels/.rels"] = TEMP_REL
	files["docProps/app.xml"] = TEMP_DOCPROPS_APP
	files["docProps/core.xml"] = TEMP_DOCPROPS_CORE
	files["word/theme/theme1.xml"] = TEMP_WORD_THEME_THEME
	files["word/styles.xml"] = f.stylesXML

	files["[Content_Types].xml"] = TEMP_CONTENT

	files["word/_rels/document.xml.rels"], err = marshal(f.DocRelation)
	if err != nil {
		return err
	}
	files["word/document.xml"], err = marshal(f.Document)
	if err != nil {
		return err
	}

	for path, data := range files {
		w, err := zipWriter.Create(path)
		if err != nil {
			return err
		}
		if _, err = w.Write([]byte(data)); err != nil {
			return err
		}
	}

	for _, m := range f.media {
		w, err := zipWriter.Create(m.path)
		if err != nil {
			return err
		}
		if _, err = w.Write(m.data); err != nil {
			return err
		}
	}

	return nil
}

func (f *File) addImageRelation(ext string, data []byte) (relID string, filename string) {
	f.imageIndex++
	filename = "image" + strconv.Itoa(f.imageIndex) + "." + ext
	rel := &RelationShip{
		ID:     "rId" + strconv.Itoa(f.rId),
		Type:   REL_IMAGE,
		Target: "media/" + filename,
	}
	f.rId++

	f.DocRelation.Relationship = append(f.DocRelation.Relationship, rel)
	f.media = append(f.media, mediaFile{
		path:        "word/media/" + filename,
		contentType: contentTypeByExt(ext),
		data:        data,
	})
	return rel.ID, filename
}

func contentTypeByExt(ext string) string {
	switch ext {
	case "png", "PNG":
		return "image/png"
	case "jpg", "JPG", "jpeg", "JPEG":
		return "image/jpeg"
	default:
		return "image/png"
	}
}

func marshal(data interface{}) (out string, err error) {
	body, err := xml.Marshal(data)
	if err != nil {
		return
	}

	out = xml.Header + string(body)
	return
}
