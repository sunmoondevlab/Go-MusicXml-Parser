package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

func TestPageMarginsL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		pmL     *PageMarginsL
		args    args
		wantErr bool
		wantObj PageMarginsL
	}{
		{
			name: "page-margins invalid decode",
			pmL:  &PageMarginsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-margins>
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
	<top-margin>85.7254</top-margin>	
	<bottom-margin>85.7255</bottom-margin>
</page-margins>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "page-margin",
					},
				},
			},
			wantErr: true,
			wantObj: PageMarginsL{},
		},
		{
			name: "page-margins invalid margin type",
			pmL:  &PageMarginsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-margins type="eve">
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
	<top-margin>85.7254</top-margin>	
	<bottom-margin>85.7255</bottom-margin>
</page-margins>
`))), start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: PageMarginsL{},
		},
		{
			name: "page-margins omit left-margin",
			pmL:  &PageMarginsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-margins>
	<right-margin>85.7253</right-margin>
	<top-margin>85.7254</top-margin>	
	<bottom-margin>85.7255</bottom-margin>
</page-margins>
`))), start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: PageMarginsL{},
		},
		{
			name: "page-margins omit right-margin",
			pmL:  &PageMarginsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-margins>
	<left-margin>85.7252</left-margin>
	<top-margin>85.7254</top-margin>	
	<bottom-margin>85.7255</bottom-margin>
</page-margins>
`))), start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: PageMarginsL{},
		},
		{
			name: "page-margins omit top-margin",
			pmL:  &PageMarginsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-margins>
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
	<bottom-margin>85.7255</bottom-margin>
</page-margins>
`))), start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: PageMarginsL{},
		},
		{
			name: "page-margins omit bottom-margin",
			pmL:  &PageMarginsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-margins>
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
	<top-margin>85.7254</top-margin>	
</page-margins>
`))), start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: PageMarginsL{},
		},
		{
			name: "page-margins",
			pmL:  &PageMarginsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-margins>
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
	<top-margin>85.7254</top-margin>	
	<bottom-margin>85.7255</bottom-margin>
</page-margins>
`))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageMarginsL{
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: nil,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
			},
		},
		{
			name: "page-margins both",
			pmL:  &PageMarginsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-margins type="both">
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
	<top-margin>85.7254</top-margin>	
	<bottom-margin>85.7255</bottom-margin>
</page-margins>
`))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageMarginsL{
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Both,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
			},
		},
		{
			name: "page-margins multiple",
			pmL:  &PageMarginsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-margins type="even">
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
	<top-margin>85.7254</top-margin>	
	<bottom-margin>85.7255</bottom-margin>
</page-margins>
<page-margins type="odd">
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
	<top-margin>85.7254</top-margin>	
	<bottom-margin>85.7255</bottom-margin>
</page-margins>
`))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageMarginsL{
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Even,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Odd,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
			},
		},
		{
			name: "page-margins empty",
			pmL:  &PageMarginsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(``))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageMarginsL{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pmL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("PageMarginsL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if diff := cmp.Diff(*tt.pmL, tt.wantObj); diff != "" {
			t.Errorf("PageMarginsL.UnmarshalXML() value is mismatch (-pmL, +wantObj):%s\n", diff)
		}
	}
}

func TestPageMarginsL_MarshalXML(t *testing.T) {
	type args struct {
		pmL *PageMarginsL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "page-margins omit left-margin",
			args: args{
				pmL: &PageMarginsL{
					{
						XMLName: xml.Name{
							Local: "page-margins",
						},
						Type: nil,
						RightMargin: &RightMargin{
							XMLName: xml.Name{
								Local: "right-margin",
							},
							RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
						},
						TopMargin: &TopMargin{
							XMLName: xml.Name{
								Local: "top-margin",
							},
							TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
						},
						BottomMargin: &BottomMargin{
							XMLName: xml.Name{
								Local: "bottom-margin",
							},
							BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
						},
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "page-margins omit right-margin",
			args: args{
				pmL: &PageMarginsL{
					{
						XMLName: xml.Name{
							Local: "page-margins",
						},
						Type: nil,
						LeftMargin: &LeftMargin{
							XMLName: xml.Name{
								Local: "left-margin",
							},
							LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
						},
						TopMargin: &TopMargin{
							XMLName: xml.Name{
								Local: "top-margin",
							},
							TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
						},
						BottomMargin: &BottomMargin{
							XMLName: xml.Name{
								Local: "bottom-margin",
							},
							BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
						},
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "page-margins omit top-margin",
			args: args{
				pmL: &PageMarginsL{
					{
						XMLName: xml.Name{
							Local: "page-margins",
						},
						Type: nil,
						LeftMargin: &LeftMargin{
							XMLName: xml.Name{
								Local: "left-margin",
							},
							LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
						},
						RightMargin: &RightMargin{
							XMLName: xml.Name{
								Local: "right-margin",
							},
							RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
						},
						BottomMargin: &BottomMargin{
							XMLName: xml.Name{
								Local: "bottom-margin",
							},
							BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
						},
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "page-margins omit bottom-margin",
			args: args{
				pmL: &PageMarginsL{
					{
						XMLName: xml.Name{
							Local: "page-margins",
						},
						Type: nil,
						LeftMargin: &LeftMargin{
							XMLName: xml.Name{
								Local: "left-margin",
							},
							LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
						},
						RightMargin: &RightMargin{
							XMLName: xml.Name{
								Local: "right-margin",
							},
							RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
						},
						TopMargin: &TopMargin{
							XMLName: xml.Name{
								Local: "top-margin",
							},
							TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
						},
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "page-margins",
			args: args{
				pmL: &PageMarginsL{
					{
						XMLName: xml.Name{
							Local: "page-margins",
						},
						Type: nil,
						LeftMargin: &LeftMargin{
							XMLName: xml.Name{
								Local: "left-margin",
							},
							LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
						},
						RightMargin: &RightMargin{
							XMLName: xml.Name{
								Local: "right-margin",
							},
							RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
						},
						TopMargin: &TopMargin{
							XMLName: xml.Name{
								Local: "top-margin",
							},
							TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
						},
						BottomMargin: &BottomMargin{
							XMLName: xml.Name{
								Local: "bottom-margin",
							},
							BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
						},
					},
				},
			},
			wantErr: false,
			wantXml: `<page-margins>
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
	<top-margin>85.7254</top-margin>
	<bottom-margin>85.7255</bottom-margin>
</page-margins>`,
		},
		{
			name: "page-margins both",
			args: args{
				pmL: &PageMarginsL{
					{
						XMLName: xml.Name{
							Local: "page-margins",
						},
						Type: &enum.MarginType.Both,
						LeftMargin: &LeftMargin{
							XMLName: xml.Name{
								Local: "left-margin",
							},
							LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
						},
						RightMargin: &RightMargin{
							XMLName: xml.Name{
								Local: "right-margin",
							},
							RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
						},
						TopMargin: &TopMargin{
							XMLName: xml.Name{
								Local: "top-margin",
							},
							TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
						},
						BottomMargin: &BottomMargin{
							XMLName: xml.Name{
								Local: "bottom-margin",
							},
							BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
						},
					},
				},
			},
			wantErr: false,
			wantXml: `<page-margins type="both">
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
	<top-margin>85.7254</top-margin>
	<bottom-margin>85.7255</bottom-margin>
</page-margins>`,
		},
		{
			name: "page-margins multiple",
			args: args{
				pmL: &PageMarginsL{
					{
						XMLName: xml.Name{
							Local: "page-margins",
						},
						Type: &enum.MarginType.Even,
						LeftMargin: &LeftMargin{
							XMLName: xml.Name{
								Local: "left-margin",
							},
							LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
						},
						RightMargin: &RightMargin{
							XMLName: xml.Name{
								Local: "right-margin",
							},
							RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
						},
						TopMargin: &TopMargin{
							XMLName: xml.Name{
								Local: "top-margin",
							},
							TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
						},
						BottomMargin: &BottomMargin{
							XMLName: xml.Name{
								Local: "bottom-margin",
							},
							BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
						},
					},
					{
						XMLName: xml.Name{
							Local: "page-margins",
						},
						Type: &enum.MarginType.Odd,
						LeftMargin: &LeftMargin{
							XMLName: xml.Name{
								Local: "left-margin",
							},
							LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
						},
						RightMargin: &RightMargin{
							XMLName: xml.Name{
								Local: "right-margin",
							},
							RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
						},
						TopMargin: &TopMargin{
							XMLName: xml.Name{
								Local: "top-margin",
							},
							TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
						},
						BottomMargin: &BottomMargin{
							XMLName: xml.Name{
								Local: "bottom-margin",
							},
							BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
						},
					},
				},
			},
			wantErr: false,
			wantXml: `<page-margins type="even">
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
	<top-margin>85.7254</top-margin>
	<bottom-margin>85.7255</bottom-margin>
</page-margins>
<page-margins type="odd">
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
	<top-margin>85.7254</top-margin>
	<bottom-margin>85.7255</bottom-margin>
</page-margins>`,
		},
		{
			name: "page-margins empty",
			args: args{
				pmL: &PageMarginsL{},
			},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.pmL, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("PageMarginsL.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("PageMarginsL.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}

func TestPageMarginsL_ValidatePageMargins(t *testing.T) {
	tests := []struct {
		name    string
		pmL     *PageMarginsL
		wantErr bool
	}{
		{
			name:    "nil",
			pmL:     nil,
			wantErr: false,
		},
		{
			name:    "empty",
			pmL:     &PageMarginsL{},
			wantErr: false,
		},
		{
			name: "len 1",
			pmL: &PageMarginsL{
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Both,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "len 3",
			pmL: &PageMarginsL{
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Both,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Both,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Both,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "duplicate key",
			pmL: &PageMarginsL{
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Both,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Both,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "not odd and not even(both,odd)",
			pmL: &PageMarginsL{
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Both,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Odd,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "not odd and not even(both(empty),even)",
			pmL: &PageMarginsL{
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: nil,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Even,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
			},
			wantErr: true,
		},
		{
			name: "odd and even",
			pmL: &PageMarginsL{
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Even,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
				{
					XMLName: xml.Name{
						Local: "page-margins",
					},
					Type: &enum.MarginType.Odd,
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
					TopMargin: &TopMargin{
						XMLName: xml.Name{
							Local: "top-margin",
						},
						TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
					},
					BottomMargin: &BottomMargin{
						XMLName: xml.Name{
							Local: "bottom-margin",
						},
						BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.pmL.ValidatePageMargins(); (err != nil) != tt.wantErr {
				t.Errorf("PageMarginsL.ValidatePageMargins() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
