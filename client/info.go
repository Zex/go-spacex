package client

import (
  "net/url"
  "fmt"
)

func (r *SpaceXClient) AboutUs(id string) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiAboutUs, id),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) AboutApi(id string) []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiAboutApi, id),
  }

  return r.Get(&url)
}
