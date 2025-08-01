package docx

const (
	TEMP_REL = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
		<Relationship Id="rId3" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties" Target="docProps/app.xml"/>
		<Relationship Id="rId2" Type="http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties" Target="docProps/core.xml"/>
		<Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="word/document.xml"/>
	</Relationships>`
	TEMP_DOCPROPS_APP  = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><Properties xmlns="http://schemas.openxmlformats.org/officeDocument/2006/extended-properties" xmlns:vt="http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes"><Application>Go DOCX</Application></Properties>`
	TEMP_DOCPROPS_CORE = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><cp:coreProperties xmlns:cp="http://schemas.openxmlformats.org/package/2006/metadata/core-properties" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:dcterms="http://purl.org/dc/terms/" xmlns:dcmitype="http://purl.org/dc/dcmitype/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"></cp:coreProperties>`
	TEMP_CONTENT       = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
		<Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>
		<Default Extension="xml" ContentType="application/xml"/>
		<Default Extension="png" ContentType="image/png"/>          
		<Default Extension="jpg" ContentType="image/jpeg"/>        
		<Default Extension="jpeg" ContentType="image/jpeg"/>   
		<Override PartName="/word/document.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml"/>
		<Override PartName="/word/styles.xml" ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml"/>
		<Override PartName="/word/theme/theme1.xml" ContentType="application/vnd.openxmlformats-officedocument.theme+xml"/>
		<Override PartName="/docProps/core.xml" ContentType="application/vnd.openxmlformats-package.core-properties+xml"/>
		<Override PartName="/docProps/app.xml" ContentType="application/vnd.openxmlformats-officedocument.extended-properties+xml"/>
	</Types>`
	TEMP_WORD_STYLE = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	<w:styles xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
		xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"
		xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"
		xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml"
		xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml"
		xmlns:w16cid="http://schemas.microsoft.com/office/word/2016/wordml/cid"
		xmlns:w16se="http://schemas.microsoft.com/office/word/2015/wordml/symex" mc:Ignorable="w14 w15 w16se w16cid">
		<w:docDefaults>
			<w:rPrDefault>
				<w:rPr>
					<w:rFonts w:asciiTheme="minorHAnsi" w:eastAsiaTheme="minorEastAsia" w:hAnsiTheme="minorHAnsi" w:cstheme="minorBidi"/>
					<w:kern w:val="2"/>
					<w:sz w:val="24"/>
					<w:szCs w:val="24"/>
					<w:lang w:val="en-US" w:eastAsia="zh-TW" w:bidi="ar-SA"/>
				</w:rPr>
			</w:rPrDefault>
			<w:pPrDefault/>
		</w:docDefaults>
		<w:style w:type="paragraph" w:default="1" w:styleId="a">
			<w:name w:val="Normal"/>
			<w:qFormat/>
			<w:rsid w:val="00A955DD"/>
			<w:pPr>
				<w:widowControl w:val="0"/>
			</w:pPr>
			<w:rPr>
				<w:szCs w:val="22"/>
			</w:rPr>
		</w:style>
		<w:style w:type="character" w:default="1" w:styleId="a0">
			<w:name w:val="Default Paragraph Font"/>
			<w:uiPriority w:val="1"/>
			<w:semiHidden/>
			<w:unhideWhenUsed/>
		</w:style>
		<w:style w:type="character" w:styleId="a1">
			<w:name w:val="Hyperlink"/>
			<w:basedOn w:val="a0"/>
			<w:uiPriority w:val="99"/>
			<w:unhideWhenUsed/>
			<w:rsid w:val="003B745A"/>
			<w:rPr>
				<w:color w:val="0563C1" w:themeColor="hyperlink"/>
				<w:u w:val="single"/>
			</w:rPr>
		</w:style>
	</w:styles>`
	TEMP_WORD_THEME_THEME = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	<a:theme xmlns:a="http://schemas.openxmlformats.org/drawingml/2006/main" name="Office-Design">
		<a:themeElements>
			<a:clrScheme name="Office">
				<a:dk1>
					<a:sysClr val="windowText" lastClr="000000"/>
				</a:dk1>
				<a:lt1>
					<a:sysClr val="window" lastClr="FFFFFF"/>
				</a:lt1>
				<a:dk2>
					<a:srgbClr val="1F497D"/>
				</a:dk2>
				<a:lt2>
					<a:srgbClr val="EEECE1"/>
				</a:lt2>
				<a:accent1>
					<a:srgbClr val="4F81BD"/>
				</a:accent1>
				<a:accent2>
					<a:srgbClr val="C0504D"/>
				</a:accent2>
				<a:accent3>
					<a:srgbClr val="9BBB59"/>
				</a:accent3>
				<a:accent4>
					<a:srgbClr val="8064A2"/>
				</a:accent4>
				<a:accent5>
					<a:srgbClr val="4BACC6"/>
				</a:accent5>
				<a:accent6>
					<a:srgbClr val="F79646"/>
				</a:accent6>
				<a:hlink>
					<a:srgbClr val="0000FF"/>
				</a:hlink>
				<a:folHlink>
					<a:srgbClr val="800080"/>
				</a:folHlink>
			</a:clrScheme>
			<a:fontScheme name="Office">
				<a:majorFont>
					<a:latin typeface="Cambria"/>
					<a:ea typeface=""/>
					<a:cs typeface=""/>
					<a:font script="Jpan" typeface="ＭＳ Ｐゴシック"/>
					<a:font script="Hang" typeface="맑은 고딕"/>
					<a:font script="Hans" typeface="宋体"/>
					<a:font script="Hant" typeface="新細明體"/>
					<a:font script="Arab" typeface="Times New Roman"/>
					<a:font script="Hebr" typeface="Times New Roman"/>
					<a:font script="Thai" typeface="Tahoma"/>
					<a:font script="Ethi" typeface="Nyala"/>
					<a:font script="Beng" typeface="Vrinda"/>
					<a:font script="Gujr" typeface="Shruti"/>
					<a:font script="Khmr" typeface="MoolBoran"/>
					<a:font script="Knda" typeface="Tunga"/>
					<a:font script="Guru" typeface="Raavi"/>
					<a:font script="Cans" typeface="Euphemia"/>
					<a:font script="Cher" typeface="Plantagenet Cherokee"/>
					<a:font script="Yiii" typeface="Microsoft Yi Baiti"/>
					<a:font script="Tibt" typeface="Microsoft Himalaya"/>
					<a:font script="Thaa" typeface="MV Boli"/>
					<a:font script="Deva" typeface="Mangal"/>
					<a:font script="Telu" typeface="Gautami"/>
					<a:font script="Taml" typeface="Latha"/>
					<a:font script="Syrc" typeface="Estrangelo Edessa"/>
					<a:font script="Orya" typeface="Kalinga"/>
					<a:font script="Mlym" typeface="Kartika"/>
					<a:font script="Laoo" typeface="DokChampa"/>
					<a:font script="Sinh" typeface="Iskoola Pota"/>
					<a:font script="Mong" typeface="Mongolian Baiti"/>
					<a:font script="Viet" typeface="Times New Roman"/>
					<a:font script="Uigh" typeface="Microsoft Uighur"/>
					<a:font script="Geor" typeface="Sylfaen"/>
				</a:majorFont>
				<a:minorFont>
					<a:latin typeface="Arial"/>
					<a:ea typeface=""/>
					<a:cs typeface=""/>
					<a:font script="Jpan" typeface="ＭＳ Ｐゴシック"/>
					<a:font script="Hang" typeface="맑은 고딕"/>
					<a:font script="Hans" typeface="宋体"/>
					<a:font script="Hant" typeface="新細明體"/>
					<a:font script="Arab" typeface="Arial"/>
					<a:font script="Hebr" typeface="Arial"/>
					<a:font script="Thai" typeface="Tahoma"/>
					<a:font script="Ethi" typeface="Nyala"/>
					<a:font script="Beng" typeface="Vrinda"/>
					<a:font script="Gujr" typeface="Shruti"/>
					<a:font script="Khmr" typeface="DaunPenh"/>
					<a:font script="Knda" typeface="Tunga"/>
					<a:font script="Guru" typeface="Raavi"/>
					<a:font script="Cans" typeface="Euphemia"/>
					<a:font script="Cher" typeface="Plantagenet Cherokee"/>
					<a:font script="Yiii" typeface="Microsoft Yi Baiti"/>
					<a:font script="Tibt" typeface="Microsoft Himalaya"/>
					<a:font script="Thaa" typeface="MV Boli"/>
					<a:font script="Deva" typeface="Mangal"/>
					<a:font script="Telu" typeface="Gautami"/>
					<a:font script="Taml" typeface="Latha"/>
					<a:font script="Syrc" typeface="Estrangelo Edessa"/>
					<a:font script="Orya" typeface="Kalinga"/>
					<a:font script="Mlym" typeface="Kartika"/>
					<a:font script="Laoo" typeface="DokChampa"/>
					<a:font script="Sinh" typeface="Iskoola Pota"/>
					<a:font script="Mong" typeface="Mongolian Baiti"/>
					<a:font script="Viet" typeface="Arial"/>
					<a:font script="Uigh" typeface="Microsoft Uighur"/>
					<a:font script="Geor" typeface="Sylfaen"/>
				</a:minorFont>
			</a:fontScheme>
			<a:fmtScheme name="Office">
				<a:fillStyleLst>
					<a:solidFill>
						<a:schemeClr val="phClr"/>
					</a:solidFill>
					<a:gradFill rotWithShape="1">
						<a:gsLst>
							<a:gs pos="0">
								<a:schemeClr val="phClr">
									<a:tint val="50000"/>
									<a:satMod val="300000"/>
								</a:schemeClr>
							</a:gs>
							<a:gs pos="35000">
								<a:schemeClr val="phClr">
									<a:tint val="37000"/>
									<a:satMod val="300000"/>
								</a:schemeClr>
							</a:gs>
							<a:gs pos="100000">
								<a:schemeClr val="phClr">
									<a:tint val="15000"/>
									<a:satMod val="350000"/>
								</a:schemeClr>
							</a:gs>
						</a:gsLst>
						<a:lin ang="16200000" scaled="1"/>
					</a:gradFill>
					<a:gradFill rotWithShape="1">
						<a:gsLst>
							<a:gs pos="0">
								<a:schemeClr val="phClr">
									<a:tint val="100000"/>
									<a:shade val="100000"/>
									<a:satMod val="130000"/>
								</a:schemeClr>
							</a:gs>
							<a:gs pos="100000">
								<a:schemeClr val="phClr">
									<a:tint val="50000"/>
									<a:shade val="100000"/>
									<a:satMod val="350000"/>
								</a:schemeClr>
							</a:gs>
						</a:gsLst>
						<a:lin ang="16200000" scaled="0"/>
					</a:gradFill>
				</a:fillStyleLst>
				<a:lnStyleLst>
					<a:ln w="9525" cap="flat" cmpd="sng" algn="ctr">
						<a:solidFill>
							<a:schemeClr val="phClr">
								<a:shade val="95000"/>
								<a:satMod val="105000"/>
							</a:schemeClr>
						</a:solidFill>
						<a:prstDash val="solid"/>
					</a:ln>
					<a:ln w="25400" cap="flat" cmpd="sng" algn="ctr">
						<a:solidFill>
							<a:schemeClr val="phClr"/>
						</a:solidFill>
						<a:prstDash val="solid"/>
					</a:ln>
					<a:ln w="38100" cap="flat" cmpd="sng" algn="ctr">
						<a:solidFill>
							<a:schemeClr val="phClr"/>
						</a:solidFill>
						<a:prstDash val="solid"/>
					</a:ln>
				</a:lnStyleLst>
				<a:effectStyleLst>
					<a:effectStyle>
						<a:effectLst>
							<a:outerShdw blurRad="40000" dist="20000" dir="5400000" rotWithShape="0">
								<a:srgbClr val="000000">
									<a:alpha val="38000"/>
								</a:srgbClr>
							</a:outerShdw>
						</a:effectLst>
					</a:effectStyle>
					<a:effectStyle>
						<a:effectLst>
							<a:outerShdw blurRad="40000" dist="23000" dir="5400000" rotWithShape="0">
								<a:srgbClr val="000000">
									<a:alpha val="35000"/>
								</a:srgbClr>
							</a:outerShdw>
						</a:effectLst>
					</a:effectStyle>
					<a:effectStyle>
						<a:effectLst>
							<a:outerShdw blurRad="40000" dist="23000" dir="5400000" rotWithShape="0">
								<a:srgbClr val="000000">
									<a:alpha val="35000"/>
								</a:srgbClr>
							</a:outerShdw>
						</a:effectLst>
						<a:scene3d>
							<a:camera prst="orthographicFront">
								<a:rot lat="0" lon="0" rev="0"/>
							</a:camera>
							<a:lightRig rig="threePt" dir="t">
								<a:rot lat="0" lon="0" rev="1200000"/>
							</a:lightRig>
						</a:scene3d>
						<a:sp3d>
							<a:bevelT w="63500" h="25400"/>
						</a:sp3d>
					</a:effectStyle>
				</a:effectStyleLst>
				<a:bgFillStyleLst>
					<a:solidFill>
						<a:schemeClr val="phClr"/>
					</a:solidFill>
					<a:gradFill rotWithShape="1">
						<a:gsLst>
							<a:gs pos="0">
								<a:schemeClr val="phClr">
									<a:tint val="40000"/>
									<a:satMod val="350000"/>
								</a:schemeClr>
							</a:gs>
							<a:gs pos="40000">
								<a:schemeClr val="phClr">
									<a:tint val="45000"/>
									<a:shade val="99000"/>
									<a:satMod val="350000"/>
								</a:schemeClr>
							</a:gs>
							<a:gs pos="100000">
								<a:schemeClr val="phClr">
									<a:shade val="20000"/>
									<a:satMod val="255000"/>
								</a:schemeClr>
							</a:gs>
						</a:gsLst>
						<a:path path="circle">
							<a:fillToRect l="50000" t="-80000" r="50000" b="180000"/>
						</a:path>
					</a:gradFill>
					<a:gradFill rotWithShape="1">
						<a:gsLst>
							<a:gs pos="0">
								<a:schemeClr val="phClr">
									<a:tint val="80000"/>
									<a:satMod val="300000"/>
								</a:schemeClr>
							</a:gs>
							<a:gs pos="100000">
								<a:schemeClr val="phClr">
									<a:shade val="30000"/>
									<a:satMod val="200000"/>
								</a:schemeClr>
							</a:gs>
						</a:gsLst>
						<a:path path="circle">
							<a:fillToRect l="50000" t="50000" r="50000" b="50000"/>
						</a:path>
					</a:gradFill>
				</a:bgFillStyleLst>
			</a:fmtScheme>
		</a:themeElements>
		<a:objectDefaults>
			<a:spDef>
				<a:spPr/>
				<a:bodyPr/>
				<a:lstStyle/>
				<a:style>
					<a:lnRef idx="1">
						<a:schemeClr val="accent1"/>
					</a:lnRef>
					<a:fillRef idx="3">
						<a:schemeClr val="accent1"/>
					</a:fillRef>
					<a:effectRef idx="2">
						<a:schemeClr val="accent1"/>
					</a:effectRef>
					<a:fontRef idx="minor">
						<a:schemeClr val="lt1"/>
					</a:fontRef>
				</a:style>
			</a:spDef>
			<a:lnDef>
				<a:spPr/>
				<a:bodyPr/>
				<a:lstStyle/>
				<a:style>
					<a:lnRef idx="2">
						<a:schemeClr val="accent1"/>
					</a:lnRef>
					<a:fillRef idx="0">
						<a:schemeClr val="accent1"/>
					</a:fillRef>
					<a:effectRef idx="1">
						<a:schemeClr val="accent1"/>
					</a:effectRef>
					<a:fontRef idx="minor">
						<a:schemeClr val="tx1"/>
					</a:fontRef>
				</a:style>
			</a:lnDef>
		</a:objectDefaults>
		<a:extraClrSchemeLst/>
	</a:theme>`
	TEMP_WORD_STYLE_TNR = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
	<w:styles xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
  	xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"
  	xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"
  	xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml"
  	xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml"
  	xmlns:w16cid="http://schemas.microsoft.com/office/word/2016/wordml/cid"
  	xmlns:w16se="http://schemas.microsoft.com/office/word/2015/wordml/symex" mc:Ignorable="w14 w15 w16se w16cid">
  	<w:docDefaults>
    	<w:rPrDefault>
      		<w:rPr>
        	<w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" w:cs="Times New Roman" w:eastAsia="Times New Roman"/>
        	<w:sz w:val="24"/>
        	<w:szCs w:val="24"/>
        	<w:lang w:val="ru-RU" w:eastAsia="zh-TW" w:bidi="ar-SA"/>
      	</w:rPr>
    	</w:rPrDefault>
    	<w:pPrDefault/>
  		</w:docDefaults>
  		<w:style w:type="paragraph" w:default="1" w:styleId="a">
    	<w:name w:val="Normal"/>
    <w:qFormat/>
    <w:pPr><w:widowControl w:val="0"/></w:pPr>
    <w:rPr>
      <w:rFonts w:ascii="Times New Roman" w:hAnsi="Times New Roman" w:cs="Times New Roman" w:eastAsia="Times New Roman"/>
      <w:szCs w:val="24"/>
    </w:rPr>
  </w:style>
  <w:style w:type="character" w:default="1" w:styleId="a0">
    <w:name w:val="Default Paragraph Font"/>
    <w:uiPriority w:val="1"/><w:semiHidden/><w:unhideWhenUsed/>
  </w:style>
  <w:style w:type="character" w:styleId="a1">
    <w:name w:val="Hyperlink"/>
    <w:basedOn w:val="a0"/><w:uiPriority w:val="99"/><w:unhideWhenUsed/>
    <w:rPr><w:color w:val="0563C1" w:themeColor="hyperlink"/><w:u w:val="single"/></w:rPr>
  </w:style>
	</w:styles>`
)
