package client

import (
  "net/url"
  "fmt"
)

const(
  ApiListCores = "/v3/cores"
  ApiGetCore = "/v3/cores/%s"
  ApiUpcomingCores = "/v3/cores/upcoming"
  ApiPastCores = "/v3/cores/past"
)

func (r *SpaceXClient) ListCores(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiListCores),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) GetCore(id string) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiGetCore, id),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) PastCores(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiPastCores),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) UpcomingCores(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiUpcomingCores),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}
