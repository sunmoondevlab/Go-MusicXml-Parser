package datatype

import (
	"fmt"
	"strconv"
	"strings"
)

type StaffNumber struct {
	Val *uint64
	fmt.Stringer
}

func (sn *StaffNumber) String() string {
	if sn == nil {
		return ""
	}
	if sn.Val == nil {
		return ""
	}
	return fmt.Sprintf("%d", *sn.Val)
}

func (sn *StaffNumber) StringPtr() *string {
	if sn == nil {
		return nil
	}
	if sn.Val == nil {
		return nil
	}
	ns := sn.String()
	return &ns
}

func (sn *StaffNumber) Value() *uint64 {
	if sn == nil {
		return nil
	}
	return sn.Val
}

func NewStaffNumber(value string) (*StaffNumber, error) {
	if len(strings.TrimSpace(value)) == 0 {
		return nil, nil
	}
	uv, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return nil, err
	}

	return &StaffNumber{
		Val: &uv,
	}, nil
}
