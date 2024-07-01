package valuetext

import (
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

type NonNegativeDecimal string

func (nnd *NonNegativeDecimal) UnmarshalText(b []byte) error {
	s := string(b)

	_, err := datatype.NewNonNegativeDecimalFromString(s)
	if err != nil {
		return fmt.Errorf("invalid value: invalid non nevative decimal, %s", err.Error())
	}
	*nnd = (NonNegativeDecimal)(s)

	return nil
}

func (nnd NonNegativeDecimal) ToEntityData() datatype.NonNegativeDecimal {
	if nnd == "" {
		return datatype.NonNegativeDecimal{}
	}

	nndV, err := datatype.NewNonNegativeDecimalFromString((string)(nnd))
	if err != nil {
		return datatype.NonNegativeDecimal{}
	}
	return *nndV
}
