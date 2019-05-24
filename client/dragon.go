package client

import (
  "net/url"
  "fmt"
)

const(
  ApiListDragons = "/v3/dragons"
  ApiGetDragon = "/v3/dragon/%s"
)

func (r *SpaceXClient) ListDragons(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiListDragons),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) GetDragon(id string) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiGetDragon, id),
  }

  return r.Get(&url)
}
