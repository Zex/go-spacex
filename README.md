# go-spacex
SpaceX Client for Go

## How to get
```
go get github.com/zex/go-spacex/client
```

## How to call
```
import (
  "github.com/zex/go-spacex/client"
)

cli := client.NewSpaceXClient()
cli.ListRockets(nil)
```
