package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/attribute"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/valuetext"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type NoteSizeL []NoteSize
type NoteSize struct {
	XMLName  xml.Name                    `xml:"note-size"`
	Type     *enum.NoteSizeTypeEnum      `xml:"type,attr,omitempty"`
	NoteSize datatype.NonNegativeDecimal `xml:",chardata"`
}

type NoteSizeRawL []NoteSizeRaw
type NoteSizeRaw struct {
	XMLName  xml.Name                     `xml:"note-size"`
	Type     *attribute.NoteSizeType      `xml:"type,attr,omitempty"`
	NoteSize valuetext.NonNegativeDecimal `xml:",chardata"`
}

func (nsL *NoteSizeL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		nsR := &NoteSizeRaw{}
		err := d.DecodeElement(nsR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			if strings.HasPrefix(err.Error(), "invalid attr:") {
				return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, nsR.XMLName.Space, nsR.XMLName.Local, err.Error())
			} else if strings.HasPrefix(err.Error(), "invalid value:") {
				return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, nsR.XMLName.Space, nsR.XMLName.Local, err.Error())
			} else {
				return err
			}
		}
		if nsR.Type == nil {
			return fmt.Errorf("attribute type is required in <note-size/>")
		}

		nsE := NoteSize{
			XMLName:  nsR.XMLName,
			Type:     (*enum.NoteSizeTypeEnum)(nsR.Type),
			NoteSize: nsR.NoteSize.ToEntityData(),
		}
		*nsL = append(*nsL, nsE)
	}

	return nil
}

func (nsL *NoteSizeL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	nsRL := NoteSizeRawL{}
	for _, nsE := range *nsL {
		if nsE.Type == nil {
			return fmt.Errorf("attribute type is required in <note-size/>")
		}

		nsR := NoteSizeRaw{
			XMLName:  nsE.XMLName,
			Type:     (*attribute.NoteSizeType)(nsE.Type),
			NoteSize: valuetext.NonNegativeDecimal(nsE.NoteSize.String()),
		}
		nsRL = append(nsRL, nsR)
		start.Name = nsE.XMLName
	}

	return e.EncodeElement(nsRL, start)
}
