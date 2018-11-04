package steamdotaapi

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"
  "net/url"
)

func TestUnsuccessfullMakeRequest(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusForbidden)
  }))
  defer ts.Close()

  api := DotaApi{"test"}
  var i interface{}
  err := api.MakeRequest(ts.URL, url.Values{}, &i)

  if err == nil {
    t.Fatalf("should return error")
  }

  _, ok := err.(HttpError)

  if ok != true {
    t.Fatalf("wrong type of error %s", err)
  }

  if i != nil {
    t.Fatalf("should not assign result")
  }
}


func TestSuccessfullMakeRequest(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{ "success": true }`)
  }))
  defer ts.Close()

  api := DotaApi{"test"}
  var i interface{}
  err := api.MakeRequest(ts.URL, url.Values{}, &i)

  if err != nil {
    t.Fatalf("returned error %s", err)
  }

  if i == nil {
    t.Fatalf("didn't assign result")
  }
}
