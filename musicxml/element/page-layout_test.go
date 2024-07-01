package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

func TestPageLayout_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName     xml.Name
		PageHeight  *PageHeight
		PageWidth   *PageWidth
		PageMargins *PageMarginsL
	}
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantObj PageLayout
	}{
		{
			name:   "page-layout invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-layout>
	<page-height>1696.94</page-height>
	<page-width>1200.48</page-width>
	<page-margins type="even">
		<left-margin>85.7252</left-margin>
		<right-margin>85.7252</right-margin>
		<top-margin>85.7252</top-margin>
		<bottom-margin>85.7252</bottom-margin>
	</page-margins>
	<page-margins type="odd">
		<left-margin>85.7252</left-margin>
		<right-margin>85.7252</right-margin>
		<top-margin>85.7252</top-margin>
		<bottom-margin>85.7252</bottom-margin>
	</page-margins>
</page-layout>`))),
				start: xml.StartElement{Name: xml.Name{Local: "page-layou"}, Attr: []xml.Attr{}},
			},
			wantErr: true,
			wantObj: PageLayout{},
		},
		{
			name:   "page-layout omit page-height",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-layout>
	<page-width>1200.48</page-width>
	<page-margins type="even">
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
</page-layout>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: PageLayout{},
		},
		{
			name:   "page-layout omit page-width",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-layout>
	<page-height>1696.94</page-height>
	<page-margins type="even">
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
</page-layout>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: PageLayout{},
		},
		{
			name:   "page-layout invalid page-margins",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-layout>
	<page-height>1696.94</page-height>
	<page-width>1200.48</page-width>
	<page-margins type="even">
		<left-margin>85.7252</left-margin>
		<right-margin>85.7253</right-margin>
		<top-margin>85.7254</top-margin>
		<bottom-margin>85.7255</bottom-margin>
	</page-margins>
	<page-margins type="both">
		<left-margin>85.7252</left-margin>
		<right-margin>85.7253</right-margin>
		<top-margin>85.7254</top-margin>
		<bottom-margin>85.7255</bottom-margin>
	</page-margins>
</page-layout>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: PageLayout{},
		},
		{
			name:   "page-layout",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-layout>
	<page-height>1696.94</page-height>
	<page-width>1200.48</page-width>
	<page-margins type="even">
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
</page-layout>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageLayout{
				XMLName: xml.Name{Local: "page-layout"},
				PageHeight: &PageHeight{
					XMLName: xml.Name{
						Local: "page-height",
					},
					PageHeight: *datatype.NewTenthsFromFloat(1696.94),
				},
				PageWidth: &PageWidth{
					XMLName: xml.Name{
						Local: "page-width",
					},
					PageWidth: *datatype.NewTenthsFromFloat(1200.48),
				},
				PageMargins: &PageMarginsL{
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
		},
		{
			name:   "page-layout omit page-height page-width",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-layout>
	<page-margins type="both">
		<left-margin>85.7252</left-margin>
		<right-margin>85.7253</right-margin>
		<top-margin>85.7254</top-margin>
		<bottom-margin>85.7255</bottom-margin>
	</page-margins>
</page-layout>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageLayout{
				XMLName: xml.Name{Local: "page-layout"},
				PageMargins: &PageMarginsL{
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
		},
		{
			name:   "page-layout omit page-height page-width margin type omit",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-layout>
	<page-margins>
		<left-margin>85.7252</left-margin>
		<right-margin>85.7253</right-margin>
		<top-margin>85.7254</top-margin>
		<bottom-margin>85.7255</bottom-margin>
	</page-margins>
</page-layout>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageLayout{
				XMLName: xml.Name{Local: "page-layout"},
				PageMargins: &PageMarginsL{
					{
						XMLName: xml.Name{
							Local: "page-margins",
						},
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
		},
		{
			name:   "page-layout omit all children",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-layout>
    </page-layout>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageLayout{
				XMLName: xml.Name{
					Local: "page-layout",
				},
			},
		},
		{
			name:   "page-layout empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageLayout{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plO := &PageLayout{
				XMLName:     tt.fields.XMLName,
				PageHeight:  tt.fields.PageHeight,
				PageWidth:   tt.fields.PageWidth,
				PageMargins: tt.fields.PageMargins,
			}
			if err := plO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("PageLayout.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(plO, &tt.wantObj); diff != "" {
				t.Errorf("PageLayout.UnmarshalXML() value is mismatch (-plO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestPageLayout_MarshalXML(t *testing.T) {
	type args struct {
		plO *PageLayout
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "page-layout omit page-height",
			args: args{
				plO: &PageLayout{
					XMLName: xml.Name{Local: "page-layout"},
					PageWidth: &PageWidth{
						XMLName: xml.Name{
							Local: "page-width",
						},
						PageWidth: *datatype.NewTenthsFromFloat(1200.48),
					},
					PageMargins: &PageMarginsL{
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
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "page-layout omit page-width",
			args: args{
				plO: &PageLayout{
					XMLName: xml.Name{Local: "page-layout"},
					PageHeight: &PageHeight{
						XMLName: xml.Name{
							Local: "page-height",
						},
						PageHeight: *datatype.NewTenthsFromFloat(1696.94),
					},
					PageMargins: &PageMarginsL{
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
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "page-layout invalid page-margins",
			args: args{
				plO: &PageLayout{
					XMLName: xml.Name{Local: "page-layout"},
					PageHeight: &PageHeight{
						XMLName: xml.Name{
							Local: "page-height",
						},
						PageHeight: *datatype.NewTenthsFromFloat(1696.94),
					},
					PageWidth: &PageWidth{
						XMLName: xml.Name{
							Local: "page-width",
						},
						PageWidth: *datatype.NewTenthsFromFloat(1200.48),
					},
					PageMargins: &PageMarginsL{
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
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "page-layout",
			args: args{
				plO: &PageLayout{
					XMLName: xml.Name{Local: "page-layout"},
					PageHeight: &PageHeight{
						XMLName: xml.Name{
							Local: "page-height",
						},
						PageHeight: *datatype.NewTenthsFromFloat(1696.94),
					},
					PageWidth: &PageWidth{
						XMLName: xml.Name{
							Local: "page-width",
						},
						PageWidth: *datatype.NewTenthsFromFloat(1200.48),
					},
					PageMargins: &PageMarginsL{
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
			},
			wantErr: false,
			wantXml: `<page-layout>
	<page-height>1696.94</page-height>
	<page-width>1200.48</page-width>
	<page-margins type="even">
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
</page-layout>`,
		},
		{
			name: "page-layout omit page-height page-width",
			args: args{
				plO: &PageLayout{
					XMLName: xml.Name{Local: "page-layout"},
					PageMargins: &PageMarginsL{
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
			},
			wantErr: false,
			wantXml: `<page-layout>
	<page-margins type="both">
		<left-margin>85.7252</left-margin>
		<right-margin>85.7253</right-margin>
		<top-margin>85.7254</top-margin>
		<bottom-margin>85.7255</bottom-margin>
	</page-margins>
</page-layout>`,
		},
		{
			name: "page-layout omit page-height page-width margin type omit",
			args: args{
				plO: &PageLayout{
					XMLName: xml.Name{Local: "page-layout"},
					PageMargins: &PageMarginsL{
						{
							XMLName: xml.Name{
								Local: "page-margins",
							},
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
			},
			wantErr: false,
			wantXml: `<page-layout>
	<page-margins>
		<left-margin>85.7252</left-margin>
		<right-margin>85.7253</right-margin>
		<top-margin>85.7254</top-margin>
		<bottom-margin>85.7255</bottom-margin>
	</page-margins>
</page-layout>`,
		},
		{
			name: "page-layout omit all children",
			args: args{
				plO: &PageLayout{
					XMLName: xml.Name{
						Local: "page-layout",
					},
				},
			},
			wantErr: false,
			wantXml: `<page-layout></page-layout>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.plO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("PageLayout.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("PageLayout.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}

func TestPageLayoutRaw_ValidatePageHeightWidth(t *testing.T) {
	type fields struct {
		XMLName     xml.Name
		PageHeight  *PageHeight
		PageWidth   *PageWidth
		PageMargins *PageMarginsL
	}
	tests := []struct {
		name    string
		fields  fields
		isNil   bool
		wantErr bool
	}{
		{
			name:    "nil",
			fields:  fields{},
			isNil:   true,
			wantErr: false,
		},
		{
			name: "omit page-height",
			fields: fields{

				XMLName: xml.Name{Local: "page-layout"},
				PageHeight: &PageHeight{
					XMLName: xml.Name{
						Local: "page-height",
					},
					PageHeight: *datatype.NewTenthsFromFloat(1696.94),
				},
				PageMargins: &PageMarginsL{},
			},
			wantErr: true,
		},
		{
			name: "omit page-height",
			fields: fields{

				XMLName: xml.Name{Local: "page-layout"},
				PageWidth: &PageWidth{
					XMLName: xml.Name{
						Local: "page-width",
					},
					PageWidth: *datatype.NewTenthsFromFloat(1200.48),
				},
				PageMargins: &PageMarginsL{},
			},
			wantErr: true,
		},
		{
			name: "valid both set page-height page-width",
			fields: fields{

				XMLName: xml.Name{Local: "page-layout"},
				PageHeight: &PageHeight{
					XMLName: xml.Name{
						Local: "page-height",
					},
					PageHeight: *datatype.NewTenthsFromFloat(1696.94),
				},
				PageWidth: &PageWidth{
					XMLName: xml.Name{
						Local: "page-width",
					},
					PageWidth: *datatype.NewTenthsFromFloat(1200.48),
				},
				PageMargins: &PageMarginsL{},
			},
			wantErr: false,
		},
		{
			name: "valid both omit page-height page-width",
			fields: fields{

				XMLName:     xml.Name{Local: "page-layout"},
				PageMargins: &PageMarginsL{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plR := &PageLayoutRaw{
				XMLName:     tt.fields.XMLName,
				PageHeight:  tt.fields.PageHeight,
				PageWidth:   tt.fields.PageWidth,
				PageMargins: tt.fields.PageMargins,
			}
			if tt.isNil {
				plR = nil
			}
			if err := plR.ValidatePageHeightWidth(); (err != nil) != tt.wantErr {
				t.Errorf("PageLayoutRaw.ValidatePageHeightWidth() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
