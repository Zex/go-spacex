package main

import (
  "github.com/zex/go-spacex/client"
  "fmt"
  "net/url"
)

func main() {
  cli := client.NewSpaceXClient()

  q := url.Values{}
  q.Add("limit", "1")
  fmt.Printf("%s\n", string(cli.UpcomingCores(q)))
  fmt.Printf("%s\n", string(cli.ListCapsules(q)))
}
