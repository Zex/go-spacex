package client

import (
  "net/url"
  "fmt"
)

const(
  ApiListShips = "/v3/ships"
  ApiGetShip = "/v3/ships/%s"
)

func (r *SpaceXClient) ListShips(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiListShips),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) GetShip(id string) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiGetShip, id),
  }

  return r.Get(&url)
}
