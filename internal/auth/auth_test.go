package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	test := map[string]struct {
		input      http.Header
		wantString string
		wantError  error
	}{
		"Empty":       {input: http.Header{}, wantString: "", wantError: ErrNoAuthHeaderIncluded},
		"One element": {input: http.Header{"Authorization": []string{"ApiKey"}}, wantError: errors.New("malformed authorization header")},
		"Success":     {input: http.Header{"Authorization": []string{"ApiKey try"}}, wantString: "try", wantError: nil},
	}

	for name, tc := range test {
		t.Run(name, func(t *testing.T) {
			gotString, gotError := GetAPIKey(tc.input)
			if !reflect.DeepEqual(gotString, tc.wantString) || gotError != tc.wantError {
				t.Fatalf("string -> expected %v, got %v \n error -> expected %v, got %v", tc.wantString, gotString, tc.wantError, gotError)
			}
		})
	}

}
