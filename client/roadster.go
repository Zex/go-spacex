package client

import (
  "net/url"
  "fmt"
)

func (r *SpaceXClient) GetRoadster(id string) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiGetRoadster, id),
  }

  return r.Get(&url)
}
