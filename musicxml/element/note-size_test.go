package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

func TestNoteSizeL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		nsL     *NoteSizeL
		args    args
		wantErr bool
		wantObj NoteSizeL
	}{
		{
			name: "note-size invalid decode",
			nsL:  &NoteSizeL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<note-size type="cue">10</note-size>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "note-siz",
					},
				},
			},
			wantErr: true,
			wantObj: NoteSizeL{},
		},
		{
			name: "note-size omit type",
			nsL:  &NoteSizeL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<note-size >10</note-size>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: NoteSizeL{},
		},
		{
			name: "note-size invalid type",
			nsL:  &NoteSizeL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<note-size type="cu">G10</note-size>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: NoteSizeL{},
		},
		{
			name: "note-size invalid decimal",
			nsL:  &NoteSizeL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<note-size type="cue">G10</note-size>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: NoteSizeL{},
		},
		{
			name: "note-size negative decimal",
			nsL:  &NoteSizeL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<note-size type="cue">-10</note-size>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: NoteSizeL{},
		},
		{
			name: "note-size",
			nsL:  &NoteSizeL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<note-size type="large">10</note-size>
<note-size type="grace-cue">10</note-size>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: NoteSizeL{
				{
					XMLName: xml.Name{
						Local: "note-size",
					},
					Type:     &enum.NoteSizeType.Large,
					NoteSize: *datatype.NewNonNegativeDecimalFixedPoint(1, 1),
				},
				{
					XMLName: xml.Name{
						Local: "note-size",
					},
					Type:     &enum.NoteSizeType.GraceCue,
					NoteSize: *datatype.NewNonNegativeDecimalFixedPoint(1, 1),
				},
			},
		},
		{
			name: "note-size empty value",
			nsL:  &NoteSizeL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<note-size type="cue"></note-size>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: NoteSizeL{
				{
					XMLName: xml.Name{
						Local: "note-size",
					},
					Type: &enum.NoteSizeType.Cue,
				},
			},
		},
		{
			name: "note-size empty",
			nsL:  &NoteSizeL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: NoteSizeL{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.nsL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("NoteSizeL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.nsL, tt.wantObj); diff != "" {
				t.Errorf("NoteSizeL.UnmarshalXML() value is mismatch (-nsL, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestNoteSizeL_MarshalXML(t *testing.T) {
	type args struct {
		nsL *NoteSizeL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "note-size omit type",
			args: args{
				nsL: &NoteSizeL{
					{
						XMLName: xml.Name{
							Local: "note-size",
						},
						NoteSize: *datatype.NewNonNegativeDecimalFixedPoint(1, 1),
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "note-size",
			args: args{
				nsL: &NoteSizeL{
					{
						XMLName: xml.Name{
							Local: "note-size",
						},
						Type:     &enum.NoteSizeType.Cue,
						NoteSize: *datatype.NewNonNegativeDecimalFixedPoint(11, 0),
					},
					{
						XMLName: xml.Name{
							Local: "note-size",
						},
						Type:     &enum.NoteSizeType.GraceCue,
						NoteSize: *datatype.NewNonNegativeDecimalFixedPoint(1, 1),
					},
				},
			},
			wantErr: false,
			wantXml: `<note-size type="cue">11</note-size>
<note-size type="grace-cue">10</note-size>`,
		},
		{
			name: "note-size empty",
			args: args{
				nsL: &NoteSizeL{},
			},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.nsL, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteSize.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("NoteSize.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
