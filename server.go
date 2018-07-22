package main

import (
    "log"

	  "github.com/google/gops/agent"
    "github.com/cristianossd/unbearable-eater/api"
)

func main() {
    if err := agent.Listen(agent.Options{}); err != nil {
        log.Fatal(err)
    }
    defer agent.Close()

    e := api.Server()

    e.Logger.Fatal(e.Start(":1323"))
}
