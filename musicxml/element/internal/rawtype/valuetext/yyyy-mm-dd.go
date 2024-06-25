package valuetext

import (
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

type YyyyMmDd string

func (ymd *YyyyMmDd) UnmarshalText(b []byte) error {
	s := string(b)

	_, err := datatype.NewYyyyMmDd(s)
	if err != nil {
		return fmt.Errorf("invalid value: invalid yyyy-mm-dd, %s", err.Error())
	}
	*ymd = (YyyyMmDd)(s)

	return nil
}

func (ymd YyyyMmDd) ToEntityData() datatype.YyyyMmDd {
	if ymd == "" {
		return datatype.YyyyMmDd{}
	}

	ymdV, err := datatype.NewYyyyMmDd((string)(ymd))
	if err != nil {
		return datatype.YyyyMmDd{}
	}
	return *ymdV
}
