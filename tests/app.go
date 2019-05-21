package main

import (
  "github.com/zex/go-spacex/client"
  "fmt"
)

func main() {
  cli := client.NewSpaceXClient()

  /**
  fmt.Println(cli.ListRockets(nil))
  fmt.Println(cli.ListShips(nil))
  fmt.Printf("%s\n", string(cli.ListPayloads(nil)))
  fmt.Printf("%s\n", string(cli.ListMissions(nil)))
  fmt.Printf("%s\n", string(cli.ListLandPads(nil)))
  fmt.Printf("%s\n", string(cli.LatestLaunch()))
  fmt.Printf("%s\n", string(cli.ListDragons(nil)))
  fmt.Printf("%s\n", string(cli.ListCores(nil)))
  fmt.Printf("%s\n", string(cli.ListCapsules(nil)))
  fmt.Printf("%s\n", string(cli.GetCapsule("C101")))
  fmt.Printf("%s\n", string(cli.ListHistory(nil)))
  fmt.Printf("%s\n", string(cli.GetHistory("19")))
  fmt.Printf("%s\n", string(cli.AboutUs()))
  fmt.Printf("%s\n", string(cli.AboutApi()))
  fmt.Printf("%s\n", string(cli.GetRoadster()))
  fmt.Printf("%s\n", string(cli.UpcomingLaunches(nil)))
  fmt.Printf("%s\n", string(cli.NextLaunch()))
  */
  fmt.Printf("%s\n", string(cli.PastLaunches(nil)))
}
