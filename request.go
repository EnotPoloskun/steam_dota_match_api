package steamdotaapi

import (
  "encoding/json"
  "net/http"
  "net/url"
  "fmt"
  "io"
  "io/ioutil"
)

type HttpError struct {
  Status int
  Url string
  Body io.Reader
}

func (e HttpError) Error() string {
  body, _ := ioutil.ReadAll(e.Body)
  return fmt.Sprintf("Http error for %v with status code %v and body: %v", e.Url, e.Status, body)
}

func (api DotaApi) MakeRequest(url string, params url.Values, target interface{}) error {
  params.Add("key", api.Key)
  url += "?" + params.Encode()
  resp, err := http.Get(url)

  if err != nil {
    return err;
  }

  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    return HttpError{resp.StatusCode, url, resp.Body}
  }

  decoder := json.NewDecoder(resp.Body)

  return decoder.Decode(&target)
}
