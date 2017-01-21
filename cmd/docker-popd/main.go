package main

import (
    "fmt"
    "os"

    "github.com/mcilloni/openbaton-docker/cmd/docker-popd/cmd"
)

func main() {
    if err := cmd.RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}