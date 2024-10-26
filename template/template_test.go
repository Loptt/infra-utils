package template

import (
	"fmt"
	"reflect"
	"testing"
)

const templateSingleValue = `
field1: ${PASSWORD}
`
const resultSingleValue = `
field1: 123456
`

const templateMultipleValues = `
config:
  - password: ${PASSWORD}
  more_config: [${EMAIL}, ${DOMAIN}]
`
const resultMultipleValues = `
config:
  - password: 123456
  more_config: [carlos@urora.mx, google.com]
`

const templateNoPassword = `
config:
  more_config: [${EMAIL}, ${DOMAIN}]
`

const templateMultiplePassword = `
config:
  - password: ${PASSWORD}
  - password2: ${PASSWORD}
  - password3: ${PASSWORD}
  - password4: ${PASSWORD}
`
const resultMultiplePassword = `
config:
  - password: 123456
  - password2: 123456
  - password3: 123456
  - password4: 123456
`

func TestLoadValues(t *testing.T) {
	tests := []struct {
		description string
		t           TemplateInterface
		m           map[string]string
		want        string
		want_err    error
	}{
		{
			description: "Test generate template single value",
			t:           &Template{content: templateSingleValue},
			m: map[string]string{
				"PASSWORD": "123456",
			},
			want:     resultSingleValue,
			want_err: nil,
		},
		{
			description: "Test generate template multiple values",
			t:           &Template{content: templateMultipleValues},
			m: map[string]string{
				"PASSWORD": "123456",
				"DOMAIN":   "google.com",
				"EMAIL":    "carlos@urora.mx",
			},
			want:     resultMultipleValues,
			want_err: nil,
		},
		{
			description: "Test value not present in template",
			t:           &Template{content: templateNoPassword},
			m: map[string]string{
				"PASSWORD": "123456",
			},
			want:     "",
			want_err: fmt.Errorf("Error!"),
		},
		{
			description: "Test multiple fields for same value",
			t:           &Template{content: templateMultiplePassword},
			m: map[string]string{
				"PASSWORD": "123456",
			},
			want:     resultMultiplePassword,
			want_err: nil,
		},
	}

	for i, test := range tests {
		got, err := test.t.GenerateFile(test.m)

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
