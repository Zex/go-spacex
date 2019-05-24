package client

import (
  "net/url"
  "fmt"
)

const(
  ApiListPayloads = "/v3/payloads"
  ApiGetPayload = "/v3/payload/%s"
)

func (r *SpaceXClient) ListPayloads(q url.Values) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiListPayloads),
    RawQuery: q.Encode(),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) GetPayload(id string) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiGetPayload, id),
  }

  return r.Get(&url)
}
