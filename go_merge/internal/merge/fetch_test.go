package merge

import (
	"reflect"
	"testing"
)

func TestFetch(t *testing.T) {
	type args struct {
		token   string
		baseUrl string
		mrs     []int
		filter  map[string]string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []MRResult
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := Fetch(tt.args.token, tt.args.baseUrl, tt.args.mrs, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Fetch() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}