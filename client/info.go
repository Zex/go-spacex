package client

import (
  "net/url"
  "fmt"
)

func (r *SpaceXClient) AboutUs() []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiAboutUs),
  }

  return r.Get(&url)
}

func (r *SpaceXClient) AboutApi() []byte {
  url := url.URL{
    Scheme: ApiScheme,
    Host: ApiHost,
    Path: fmt.Sprintf(ApiAboutApi),
  }

  return r.Get(&url)
}
