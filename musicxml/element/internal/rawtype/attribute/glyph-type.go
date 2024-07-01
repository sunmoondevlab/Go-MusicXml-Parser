package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type GlyphType enum.GlyphTypeEnum

func (gt *GlyphType) UnmarshalXMLAttr(attr xml.Attr) error {
	gtV, err := enum.ToGlyphType(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*gt = (GlyphType)(*gtV)

	return nil
}
