package steamdotaapi

import (
  "fmt"
  "net/http"
  "net/http/httptest"
  "testing"
  "net/url"
)

func TestGetMatchDetails(t *testing.T) {
  ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
      {
        "result": {
          "radiant_win": true
        }
      }
    `)
  }))
  defer ts.Close()

  api := DotaApi{"test"}
  var i = MatchResultJson{}
  err := api.MakeRequest(ts.URL, url.Values{}, &i)

  if err != nil {
    t.Fatalf("returned error %s", err)
  }

  if (!i.Result.RadiantWin) {
    t.Fatalf("wrong radiant win")
  }
}

func TestBarracksStatusIsDestroyed(t *testing.T) {
  var barracksStatus BarracksStatus = 5

  if (barracksStatus.IsDestroyed(TopMelee)) {
    t.Errorf("wrong barracks status for %v and %v", barracksStatus, TopMelee)
  }

  if (barracksStatus.IsDestroyed(MiddleMelee)) {
    t.Errorf("wrong barracks status for %v and %v", barracksStatus, MiddleMelee)
  }

  if (!barracksStatus.IsDestroyed(BottomMelee)) {
    t.Errorf("wrong barracks status for %v and %v", barracksStatus, BottomMelee)
  }
}
