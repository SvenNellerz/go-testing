package signal

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatalf("http.NewRequest() err = %s", err)
	}
	Handler(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Handler() status = %d; want %d", resp.StatusCode, http.StatusOK)
	}

	contentType := resp.Header.Get("Content-Type")

	if contentType != "application/json" {
		t.Errorf("Handler() Content-Type = %q; want %q", contentType, "application/json")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll() err = %s", err)
	}

	var p Person
	err = json.Unmarshal(data, &p)
	if err != nil {
		t.Fatalf("json.Unmarshal() err = %s", err)
	}

	wantAge := 30
	if p.Age != wantAge {
		t.Errorf("person.Age = %d; want %d", p.Age, wantAge)
	}
	wantName := "Bob Jones"
	if p.Name != wantName {
		t.Errorf("person.Name = %q; want %q", p.Name, wantName)
	}
}
