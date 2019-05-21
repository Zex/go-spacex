package client

import (
  "net/http"
  "net/url"
  "log"
  "io/ioutil"
)

type SpaceXClient struct {
  client *http.Client
}

func NewSpaceXClient() *SpaceXClient {
  return &SpaceXClient{
    client: &http.Client{},
  }
}

func (r *SpaceXClient) Get(url *url.URL) []byte {
  rsp, err := r.client.Get(url.String())
  if err != nil {
    return nil
  }
  defer rsp.Body.Close()

  if rsp.StatusCode != http.StatusOK {
    log.Printf("%d: %s\n", rsp.StatusCode, rsp.Status)
    return nil
  }

  body, err := ioutil.ReadAll(rsp.Body)
  return body
}
