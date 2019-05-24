package client

import (
  "net/url"
  "fmt"
)

const(
  ApiListLandPads = "/v3/landpads"
  ApiGetLandPad = "/v3/landpads/%s"
)

const(
  ApiListLaunchPads = "/v3/launchpads"
  ApiGetLaunchPad = "/v3/launchpads/%s"
)

func (r *SpaceXClient) ListLaunchPads(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiListLaunchPads),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) GetLaunchPad(id string) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiGetLaunchPad, id),
  }

  return r.Get(&url)
}
