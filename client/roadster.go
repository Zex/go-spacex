package client

import (
  "net/url"
  "fmt"
)

func (r *SpaceXClient) GetRoadster() []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiGetRoadster),
  }

  return r.Get(&url)
}
