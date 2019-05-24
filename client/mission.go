package client

import (
  "net/url"
  "fmt"
)

const(
  ApiListMissions = "/v3/missions"
  ApiGetMission = "/v3/mission/%s"
)

func (r *SpaceXClient) ListMissions(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiListMissions),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) GetMission(id string) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiGetMission, id),
  }

  return r.Get(&url)
}
