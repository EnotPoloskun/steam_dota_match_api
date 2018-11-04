package steamdotaapi

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"
  "net/url"
)

func TestGetMatchHistory(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
      {
        "result": {
          "status": 1,
          "total_results": 100,
          "matches": [
            {
              "match_id": 123,
              "players": [
                {
                  "account_id": 321
                }
              ]
            }
          ]
        }
      }
    `)
  }))
  defer ts.Close()

  api := DotaApi{"test"}
  var i = MatchHistoryJson{}
  err := api.MakeRequest(ts.URL, url.Values{}, &i)

  if err != nil {
    t.Fatalf("returned error %s", err)
  }

  if (i.Result.Matches[0].Players[0].AccountId != 321) {
    t.Fatalf("wrong account id")
  }
}
