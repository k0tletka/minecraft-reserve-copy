package config

import (
    "fmt"
    "os"
    "github.com/BurntSushi/toml"
)

func ReadConfiguration(filepath string) (*Configuration, error) {
    cfg := getDefaultConfiguration()
    cfg.validConfigConditions = []validCondition{
        validConditionFunc(authCondition),
    }

    fileData, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }

    errors := cfg.Parse(fileData)
    if len(errors) != 0 {
        for _, err := range errors {
            fmt.Fprintf(os.Stderr, "Error occured when loading configuration: %s", err)
        }

        return nil, ErrConfigParseError
    }

    return cfg, nil
}

func getDefaultConfiguration() *Configuration {
    return &Configuration{
        Webdav: {
            WebdavHost: "http://localhost",
            UseAuth: false,
        },
    }
}

// Condition functions
func authCondition(md *toml.Metadata, conf *Configuration) error {
    return nil
}
