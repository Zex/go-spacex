package main

import (
  "go-spacex/client"
  "fmt"
)

func main() {
  cli := client.NewSpaceXClient()
  fmt.Println(cli.ListRockets(nil))
}
