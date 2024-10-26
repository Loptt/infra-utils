package filegenerator

import (
	"reflect"
	"testing"

	"github.com/Loptt/infra-utils/template"
)

func TestLoadValues(t *testing.T) {
	tests := []struct {
		description string
		fg          *FileGenerator
		path        string
		want        *FileValues
		want_err    error
	}{
		{
			description: "Test load values single entry",
			fg:          NewFileGenerator(&template.Template{}),
			path:        "./testdata/single-entry.yaml",
			want: &FileValues{Files: []File{
				{
					Name: "test.yaml",
					Values: map[string]string{
						"a": "b",
						"c": "d",
					},
				},
			}},
			want_err: nil,
		},
	}

	for i, test := range tests {
		err := test.fg.LoadValues(test.path)
		got := test.fg.fv

		// If we are exepcting an error, then check that we actually get one.
		if test.want_err != nil {
			if err == nil {
				t.Errorf("Test #%d %s: want err %v, got %v", i, test.description, test.want_err, err)
			}
		} else {
			// In this case we don't expect an error, so any error should fail the test.
			if err != nil {
				t.Errorf("Test #%d %s: found error %v", i, test.description, err)
			} else if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Test #%d %s: got \n%v,\nwant\n%v", i, test.description, got, test.want)
			}
		}
	}
}
