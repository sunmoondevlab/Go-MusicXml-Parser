package valuetext

import (
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

type Millimeters string

func (m *Millimeters) UnmarshalText(b []byte) error {
	s := string(b)

	_, err := datatype.NewMillimetersFromString(s)
	if err != nil {
		return fmt.Errorf("invalid value: invalid Millimeters, %s", err.Error())
	}
	*m = (Millimeters)(s)

	return nil
}

func (m Millimeters) ToEntityData() datatype.Millimeters {
	if m == "" {
		return datatype.Millimeters{}
	}

	mV, err := datatype.NewMillimetersFromString((string)(m))
	if err != nil {
		return datatype.Millimeters{}
	}
	return *mV
}
