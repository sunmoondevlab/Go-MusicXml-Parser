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
		Composer   CreatorElement
		Arranger   CreatorElement
		Lyricist   CreatorElement
		Poet       CreatorElement
		Translator CreatorElement
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
			fields: fields{Composer: CreatorElement{}, Arranger: CreatorElement{}, Lyricist: CreatorElement{}, Poet: CreatorElement{}, Translator: CreatorElement{}},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<creator type="composer">作曲者:composs</creator>
		    <creator type="arranger">編曲者:arran</creator>
		    <creator type="lyricist">作詞者:lilly</creator>
		    <creator type="poet">作詩者:poet</creator>
		    <creator type="translator">訳詞者:trann</creator>`))),
				start: xml.StartElement{Name: xml.Name{Space: "", Local: "creato"}, Attr: []xml.Attr{}}},
			wantErr: true,
			wantObj: Creator{
				Composer:   CreatorElement{},
				Arranger:   CreatorElement{},
				Lyricist:   CreatorElement{},
				Poet:       CreatorElement{},
				Translator: CreatorElement{},
			},
		},
		{
			name: "creator invalid type",
			fields: fields{
				Composer:   CreatorElement{},
				Arranger:   CreatorElement{},
				Lyricist:   CreatorElement{},
				Poet:       CreatorElement{},
				Translator: CreatorElement{},
			},
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
				Composer: CreatorElement{
					XMLName: xml.Name{
						Space: "",
						Local: "creator",
					},
					Type:    enum.CreatorType.Composer,
					Creator: "作曲者:composs",
				},
				Arranger: CreatorElement{
					XMLName: xml.Name{
						Space: "",
						Local: "creator",
					},
					Type:    enum.CreatorType.Arranger,
					Creator: "編曲者:arran",
				},
				Lyricist: CreatorElement{
					XMLName: xml.Name{
						Space: "",
						Local: "creator",
					},
					Type:    enum.CreatorType.Lyricist,
					Creator: "作詞者:lilly",
				},
				Poet: CreatorElement{
					XMLName: xml.Name{
						Space: "",
						Local: "creator",
					},
					Type:    enum.CreatorType.Poet,
					Creator: "作詩者:poet",
				},
				Translator: CreatorElement{
					XMLName: xml.Name{
						Space: "",
						Local: "creator",
					},
					Type:    enum.CreatorType.Translator,
					Creator: "訳詞者:trann",
				},
			},
		},
		{
			name: "creator",
			fields: fields{
				Composer:   CreatorElement{},
				Arranger:   CreatorElement{},
				Lyricist:   CreatorElement{},
				Poet:       CreatorElement{},
				Translator: CreatorElement{},
			},
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
				Composer: CreatorElement{
					XMLName: xml.Name{
						Space: "",
						Local: "creator",
					},
					Type:    enum.CreatorType.Composer,
					Creator: "作曲者:composs",
				},
				Arranger: CreatorElement{
					XMLName: xml.Name{
						Space: "",
						Local: "creator",
					},
					Type:    enum.CreatorType.Arranger,
					Creator: "編曲者:arran",
				},
				Lyricist: CreatorElement{
					XMLName: xml.Name{
						Space: "",
						Local: "creator",
					},
					Type:    enum.CreatorType.Lyricist,
					Creator: "作詞者:lilly",
				},
				Poet: CreatorElement{
					XMLName: xml.Name{
						Space: "",
						Local: "creator",
					},
					Type:    enum.CreatorType.Poet,
					Creator: "作詩者:poet",
				},
				Translator: CreatorElement{
					XMLName: xml.Name{
						Space: "",
						Local: "creator",
					},
					Type:    enum.CreatorType.Translator,
					Creator: "訳詞者:trann",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Creator{
				Composer:   tt.fields.Composer,
				Arranger:   tt.fields.Arranger,
				Lyricist:   tt.fields.Lyricist,
				Poet:       tt.fields.Poet,
				Translator: tt.fields.Translator,
			}
			if err := c.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Creator.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*c, tt.wantObj); diff != "" {
				t.Errorf("Creator.UnmarshalXML() value is mismatch (-*c +tt.wantObj):%s\n", diff)
			}
		})
	}
}
