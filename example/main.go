package main

import (
	"github.com/Zhaisan/docx"
)

func main() {
	f := docx.NewFile()

	para := f.AddParagraph()

	para.AddText("test")

	para.AddText("test font size").Size(22)
	para.AddText("test color").Color("808080")
	para.AddText("test font size and color").Size(22).Color("121212")

	nextPara := f.AddParagraph()
	nextPara.AddLink("google", `http://google.com`)

	f.Save("./test.docx")
}
