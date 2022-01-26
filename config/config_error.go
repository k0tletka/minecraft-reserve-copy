package config

import (
    "fmt"
)

var (
    ErrConfigParseError = errors.New("config parse error")
)

type ConfigOptionIsNotDefined struct {
    ConfigOptionName string
}

func (c *ConfigOptionIsNotDefined) Error() string {
    return fmt.Sprintf("option is not defined in the config: %s", c.ConfigOptionName)
}
