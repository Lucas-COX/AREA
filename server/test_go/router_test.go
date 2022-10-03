package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
func TestNew(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("[]"))
	})

	req := httptest.NewRequest("GET", "/Accept", nil)
	rr := httptest.NewRecorder()

	handler := middleware.Logger(testHandler)
	handler.ServeHTTP(rr, req)
}
*/

func TestGetID(t *testing.T) {
	tests := []struct {
		name           string
		rec            *httptest.ResponseRecorder
		req            *http.Request
		expectedBody   string
		expectedHeader string
	}{
		{
			name:         "OK_1",
			rec:          httptest.NewRecorder(),
			req:          httptest.NewRequest("GET", "/ping/1", nil),
			expectedBody: `article ID:1`,
		},
		{
			name:         "OK_100",
			rec:          httptest.NewRecorder(),
			req:          httptest.NewRequest("GET", "/ping/100", nil),
			expectedBody: `article ID:100`,
		},
		{
			name:         "BAD_REQUEST",
			rec:          httptest.NewRecorder(),
			req:          httptest.NewRequest("PUT", "/ping/bad", nil),
			expectedBody: fmt.Sprintf("%s\n", http.StatusText(http.StatusBadRequest)),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//New(http.HandlerFunc()).ServeHTTP(test.rec, test.req)
			if test.expectedBody != test.rec.Body.String() {
				t.Errorf("Got: \t\t%s\n\tExpected: \t%s\n", test.rec.Body.String(), test.expectedBody)
			}
		})
	}
}

/*
func TestNew(t *testing.T) (*http.Response, string) {
	req, err := http.NewRequest("Get", "/Accept", nil)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer resp.Body.Close()

	return resp, string(respBody)
}
*/
