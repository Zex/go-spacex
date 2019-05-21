package client

import (
  "net/url"
  "fmt"
)

func (r *SpaceXClient) ListCapsules(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiListCapsules),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) GetCapsule(id string) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiGetCapsule, id),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) PastCapsules(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiPastCapsules),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) UpcomingCapsules(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiUpcomingCapsules),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}
