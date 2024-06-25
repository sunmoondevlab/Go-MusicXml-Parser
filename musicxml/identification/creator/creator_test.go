package creator

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

func TestCreator_UnmarshalXML(t *testing.T) {
	type fields struct {
		Composer   *CreatorElement
		Arranger   *CreatorElement
		Lyricist   *CreatorElement
		Poet       *CreatorElement
		Translator *CreatorElement
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
		wantObj Creator
	}{
		{
			name:   "creator invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<creator type="composer">作曲者:composs</creator>
		    <creator type="arranger">編曲者:arran</creator>
		    <creator type="lyricist">作詞者:lilly</creator>
		    <creator type="poet">作詩者:poet</creator>
		    <creator type="translator">訳詞者:trann</creator>`))),
				start: xml.StartElement{Name: xml.Name{Local: "creato"}, Attr: []xml.Attr{}}},
			wantErr: true,
			wantObj: Creator{},
		},
		{
			name:   "creator invalid type",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<creator type="composer">作曲者:composs</creator>
    <creator type="arranger">編曲者:arran</creator>
    <creator type="lyricist">作詞者:lilly</creator>
    <creator type="poet">作詩者:poet</creator>
	<creator type="translator">訳詞者:trann</creator>
    <creator type="translato">訳詞者:trann</creator>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: Creator{
				XMLName: xml.Name{
					Local: "creator",
				},
				Composer: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Composer,
					Creator: "作曲者:composs",
				},
				Arranger: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Arranger,
					Creator: "編曲者:arran",
				},
				Lyricist: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Lyricist,
					Creator: "作詞者:lilly",
				},
				Poet: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Poet,
					Creator: "作詩者:poet",
				},
				Translator: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Translator,
					Creator: "訳詞者:trann",
				},
			},
		},
		{
			name:   "creator",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<creator type="composer">作曲者:composs</creator>
    <creator type="arranger">編曲者:arran</creator>
    <creator type="lyricist">作詞者:lilly</creator>
    <creator type="poet">作詩者:poet</creator>
    <creator type="translator">訳詞者:trann</creator>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Creator{
				XMLName: xml.Name{
					Local: "creator",
				},
				Composer: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Composer,
					Creator: "作曲者:composs",
				},
				Arranger: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Arranger,
					Creator: "編曲者:arran",
				},
				Lyricist: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Lyricist,
					Creator: "作詞者:lilly",
				},
				Poet: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Poet,
					Creator: "作詩者:poet",
				},
				Translator: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Translator,
					Creator: "訳詞者:trann",
				},
			},
		},
		{
			name:   "creator omit pattern 1",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`
    <creator type="arranger">編曲者:arran</creator>
    <creator type="poet">作詩者:poet</creator>
    `))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Creator{
				XMLName: xml.Name{
					Local: "creator",
				},
				Arranger: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Arranger,
					Creator: "編曲者:arran",
				},
				Poet: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Poet,
					Creator: "作詩者:poet",
				},
			},
		},
		{
			name:   "creator omit pattern 2",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<creator type="composer">作曲者:composs</creator>
    <creator type="lyricist">作詞者:lilly</creator>
    <creator type="translator">訳詞者:trann</creator>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Creator{
				XMLName: xml.Name{
					Local: "creator",
				},
				Composer: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Composer,
					Creator: "作曲者:composs",
				},
				Lyricist: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Lyricist,
					Creator: "作詞者:lilly",
				},
				Translator: &CreatorElement{
					XMLName: xml.Name{
						Local: "creator",
					},
					Type:    enum.CreatorType.Translator,
					Creator: "訳詞者:trann",
				},
			},
		},
		{
			name:   "creator empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Creator{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cO := &Creator{
				Composer:   tt.fields.Composer,
				Arranger:   tt.fields.Arranger,
				Lyricist:   tt.fields.Lyricist,
				Poet:       tt.fields.Poet,
				Translator: tt.fields.Translator,
			}
			if err := cO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Creator.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*cO, tt.wantObj); diff != "" {
				t.Errorf("Creator.UnmarshalXML() value is mismatch (-*cO +tt.wantObj):%s\n", diff)
			}
		})
	}
}

func TestCreator_MarshalXML(t *testing.T) {
	type args struct {
		cO *Creator
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "craetor",
			args: args{
				cO: &Creator{
					XMLName: xml.Name{
						Local: "creator",
					},
					Composer: &CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Composer,
						Creator: "作曲者:composs",
					},
					Arranger: &CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Arranger,
						Creator: "編曲者:arran",
					},
					Lyricist: &CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Lyricist,
						Creator: "作詞者:lilly",
					},
					Poet: &CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Poet,
						Creator: "作詩者:poet",
					},
					Translator: &CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Translator,
						Creator: "訳詞者:trann",
					},
				},
			},
			wantErr: false,
			wantXml: `<creator type="composer">作曲者:composs</creator>
<creator type="arranger">編曲者:arran</creator>
<creator type="lyricist">作詞者:lilly</creator>
<creator type="poet">作詩者:poet</creator>
<creator type="translator">訳詞者:trann</creator>`,
		},
		{
			name: "craetor omit pattern 1",
			args: args{
				cO: &Creator{
					XMLName: xml.Name{
						Local: "creator",
					},
					Composer: &CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Composer,
						Creator: "作曲者:composs",
					},
					Lyricist: &CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Lyricist,
						Creator: "作詞者:lilly",
					},
					Translator: &CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Translator,
						Creator: "訳詞者:trann",
					},
				},
			},
			wantErr: false,
			wantXml: `<creator type="composer">作曲者:composs</creator>
<creator type="lyricist">作詞者:lilly</creator>
<creator type="translator">訳詞者:trann</creator>`,
		},
		{
			name: "craetor omit pattern 2",
			args: args{
				cO: &Creator{
					XMLName: xml.Name{
						Local: "creator",
					},
					Arranger: &CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Arranger,
						Creator: "編曲者:arran",
					},
					Poet: &CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Poet,
						Creator: "作詩者:poet",
					},
				},
			},
			wantErr: false,
			wantXml: `<creator type="arranger">編曲者:arran</creator>
<creator type="poet">作詩者:poet</creator>`,
		},
		{
			name: "craetor empty all",
			args: args{
				cO: &Creator{},
			},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.cO, "", "  ")
			if (err != nil) != tt.wantErr {
				t.Errorf("Creator.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Creator.MarshalXML() value is mismatch (-ox +tt.wantXml):%s\n", diff)
			}
		})
	}
}
