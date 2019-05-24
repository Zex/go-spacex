package client

import (
  "net/url"
  "fmt"
)

const (
  ApiListHistory = "/v3/history"
  ApiGetHistory = "/v3/history/%s"
)

func (r *SpaceXClient) ListHistory(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiListHistory),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) GetHistory(id string) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiGetHistory, id),
  }

  return r.Get(&url)
}
