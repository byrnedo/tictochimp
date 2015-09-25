package models

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"testing"
)

const (
	TICTAIL_TEST_URL = "https://api.tictail.com"
	TICTAIL_TEST_KEY = "testToken1234"
)

func TestGetMe(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", TICTAIL_TEST_URL+"/v1/me",
		func(req *http.Request) (*http.Response, error) {
			if req.Header.Get("Authorization") == "Bearer "+TICTAIL_TEST_KEY {
				return httpmock.NewStringResponse(200, TICTAIL_MOCK_GET_ME_200_RESPONSE), nil
			}
			return httpmock.NewStringResponse(401, "{}"), nil
		},
	)

	tt := NewTictail(TICTAIL_TEST_KEY)

	_, err := tt.GetMe()
	if err != nil {
		t.Error(err.Error())
	}

	tt = NewTictail("WRONGKEY")

	_, err = tt.GetMe()
	if err == nil {
		t.Error(err.Error())
	}
}
