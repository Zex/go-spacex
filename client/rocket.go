package client

import (
  "net/url"
  "fmt"
)

const(
  ApiListRockets = "/v3/rockets"
  ApiGetRocket = "/v3/rockets/%s"
)

func (r *SpaceXClient) ListRockets(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: ApiListRockets,
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) GetRocket(id string) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiGetRocket, id),
  }

  return r.Get(&url)
}
