package config

import (
    "fmt"
    "os"
    "io"
    "github.com/BurntSushi/toml"
)

func ReadConfiguration(filepath string) (*Configuration, error) {
    cfg := GetDefaultConfiguration()
    cfg.validConfigConditions = []validCondition{
        validConditionFunc(authCondition),
    }

    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }

    fileData, err := io.ReadAll(file)
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

func GetDefaultConfiguration() *Configuration {
    return &Configuration{
        WorldPath: "world",
        ArchiveNameTemplate: "{{.WorldSaveName}}-{{.Datetime}}.zip",
        TimeTemplate: "01-02-03-04-05PM",
        Webdav: WebdavConfig{
            WebdavHost: "http://localhost",
            WebdavSavePath: "backup/",
            UseAuth: false,
        },
    }
}

// Condition functions
func authCondition(md *toml.MetaData, conf *Configuration) error {
    return nil
}
