package valuetext

import (
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

type Tenths string

func (t *Tenths) UnmarshalText(b []byte) error {
	s := string(b)

	_, err := datatype.NewTenthsFromString(s)
	if err != nil {
		return fmt.Errorf("invalid value: invalid tenths, %s", err.Error())
	}
	*t = (Tenths)(s)

	return nil
}

func (t Tenths) ToEntityData() datatype.Tenths {
	if t == "" {
		return datatype.Tenths{}
	}

	tV, err := datatype.NewTenthsFromString((string)(t))
	if err != nil {
		return datatype.Tenths{}
	}
	return *tV
}
