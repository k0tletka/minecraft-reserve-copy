package main

import (
    "flag"
)

// Flag initialization
var (
    optConfig =     flag.String("c", "configuration.toml", "Configuration file to load")
    optWorldPath =  flag.String("p", "world", "Path to minecraft save to reserve")
    optLogFile =    flag.String("o", "", "Path for file to log. If empty, log to stdout")
)

func main() {
}
