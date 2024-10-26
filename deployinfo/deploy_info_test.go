package deployinfo

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Loptt/infra-utils/file"
)

func TestNewDeployInfo(t *testing.T) {
	tests := []struct {
		description    string
		deployInfoPath string
		want           *DeployInfo
		want_err       error
	}{
		{
			description:    "Test create simple deploy info one prod",
			deployInfoPath: "./testdata/simple-prod-deploy-info.yaml",
			want:           &DeployInfo{DeployData{ProdLocations: []string{"/production/mach/website"}}},
			want_err:       nil,
		},
		{
			description:    "Test create simple deploy info one staging",
			deployInfoPath: "./testdata/simple-staging-deploy-info.yaml",
			want:           &DeployInfo{DeployData{StagingLocations: []string{"/production/mach/website"}}},
			want_err:       nil,
		},
		{
			description:    "Test create deploy info with multiple locations",
			deployInfoPath: "./testdata/multi-deploy-info.yaml",
			want: &DeployInfo{
				DeployData{
					ProdLocations: []string{
						"/production/mach/website",
						"/production/mach3/website3",
						"/otherloc/qwerty/product5",
					},
					StagingLocations: []string{
						"/production/mach4/db",
						"/production/mach500/website3",
						"/otherloc/asdf/product10",
					},
				}},
			want_err: nil,
		},
		{
			description:    "Test create with invalid deploy info format",
			deployInfoPath: "./testdata/bad-deploy-info.yaml",
			want:           nil,
			want_err:       errors.New("Error!"),
		},
	}

	for i, test := range tests {
		f := &file.FileManager{}
		got, err := NewDeployInfo(test.deployInfoPath, f)

		// If we are exepcting an error, then check that we actually get one.
		if test.want_err != nil {
			if err == nil {
				t.Errorf("Test #%d %s: want err %v, got %v", i, test.description, test.want_err, err)
			}
		} else {
			// In this case we don't expect an error, so any error should fail the test.
			if err != nil {
				t.Errorf("Test #%d %s: found error %v", i, test.description, err)
			} else if !reflect.DeepEqual(*got, *test.want) {
				t.Errorf("Test #%d %s: got \n%v,\nwant\n%v", i, test.description, got, test.want)
			}
		}
	}
}
