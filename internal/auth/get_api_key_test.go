package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_api_key_good(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey beeebooop")
	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf(err.Error())
	}
	want := "beeebooop"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func Test_api_key_malformed(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey")
	_, err := GetAPIKey(headers)
	if err.Error() != "malformed authorization header" {
		t.Fatalf("want 'malformed authorization header' err, got %v", err)
	}
}

func Test_no_auth_header(t *testing.T) {
	headers := http.Header{}
	headers.Set("Bearer", "Token")
	_, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("want ErrNoAuthHeaderIncluded err, got %v", err)
	}
}
