package main

import (
  "github.com/zex/go-spacex/client"
  "fmt"
)

func main() {
  cli := client.NewSpaceXClient()

  fmt.Println(cli.ListRockets(nil))
  fmt.Println(cli.ListShips(nil))
  fmt.Printf("%s", string(cli.ListPayloads(nil)))
  fmt.Printf("%s", string(cli.ListMissions(nil)))
  fmt.Printf("%s", string(cli.ListLandPads(nil)))
  fmt.Printf("%s", string(cli.LatestLaunch()))
  fmt.Println("%s", string(cli.ListDragons(nil)))
}
