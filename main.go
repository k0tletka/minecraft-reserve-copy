package main

import (
    "flag"
    "fmt"
    "os"

    "minecraft-reverse-copy/config"
)

const (
    // Exit codes
    GeneralIssue = 1
    ConfigurationReadError = 2
)

// Flag initialization
var (
    optConfig = flag.String("c", "", "Configuration file to load")
)

func main() {
    var (
        conf *config.Configuration
        err error
    )

    flag.Parse()

    if *optConfig == "" {
        // Load default configuration
        conf = config.GetDefaultConfiguration()
    } else {
        if conf, err = config.ReadConfiguration(*optConfig); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(ConfigurationReadError)
        }
    }

    // Start zip process
    fmt.Println(conf)
}
