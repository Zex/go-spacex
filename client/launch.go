package client

import (
  "net/url"
  "fmt"
)

const(
  ApiListLaunches = "/v3/launches"
  ApiGetLaunch = "/v3/launches/%s"
  ApiPastLaunches = "/v3/launches/past"
  ApiUpcomingLaunches = "/v3/launches/upcoming"
  ApiLatestLaunch = "/v3/launches/latest"
  ApiNextLaunch = "/v3/launches/next"
)

func (r *SpaceXClient) ListLaunches(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiListLaunches),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) GetLaunch(id string) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiGetLaunch, id),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) PastLaunches(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiPastLaunches),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) UpcomingLaunches(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiUpcomingLaunches),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) LatestLaunch() []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiLatestLaunch),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) NextLaunch() []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiNextLaunch),
  }

  return r.Get(&url)
}
