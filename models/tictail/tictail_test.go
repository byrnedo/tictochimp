package tictail

import (
	"encoding/json"
	"fmt"
	"github.com/byrnedo/tictochimp/models/tictail/spec"
	"github.com/jarcoal/httpmock"
	"net/http"
	"reflect"
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
	tt := NewTictail("WRONGKEY")

	response, err := tt.GetMe()
	if err == nil {
		t.Error(err.Error())
	}

	tt = NewTictail(TICTAIL_TEST_KEY)

	response, err = tt.GetMe()
	if err != nil {
		t.Error(err.Error())
	} else if response == nil {
		t.Error("Nil response")
	}

	expectedData := &spec.MeResponse{}
	err = json.Unmarshal([]byte(TICTAIL_MOCK_GET_ME_200_RESPONSE), expectedData)
	if err != nil {
		t.Error(err.Error())
	}

	if fmt.Sprintf("%#v", response) != fmt.Sprintf("%#v", expectedData) {
		t.Error("Response and expected data don't match")
	}
}

func TestGetOrders(t *testing.T) {

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", TICTAIL_TEST_URL+"/v1/store/x1234/orders",
		func(req *http.Request) (*http.Response, error) {
			if req.Header.Get("Authorization") == "Bearer "+TICTAIL_TEST_KEY {
				return httpmock.NewStringResponse(200, TICTAIL_MOCK_GET_ORDERS_200_RESPONSE), nil
			}
			return httpmock.NewStringResponse(401, "{}"), nil
		},
	)

	tt := NewTictail(TICTAIL_TEST_KEY)

	response, err := tt.GetAllOrders("x1234")
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("response: %#v", response)

	var expectedData []spec.OrdersResponse
	err = json.Unmarshal([]byte(TICTAIL_MOCK_GET_ORDERS_200_RESPONSE), &expectedData)
	if err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(expectedData, response) {
		t.Error("Response and mock data didn't match")
	}

}
