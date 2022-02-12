package config

import (
    "github.com/BurntSushi/toml"
)

type Configuration struct {
    Webdav      WebdavConfig `toml:"webdav"`
    LogFile     string `toml:"log_file"`
    WorldPath   string `toml:"world_path"`

    validConfigConditions []validCondition
    parseMetadata toml.MetaData
}

func (c *Configuration) Parse(data []byte) (errors []error) {
    var err error
    c.parseMetadata, err = toml.Decode(string(data), c)

    if err != nil {
        errors = append(errors, err)
        return
    }

    errors = c.checkConditions()
    return
}

func (c *Configuration) checkConditions() []error {
    resultErrors := make([]error, 0, len(c.validConfigConditions))

    for _, condition := range c.validConfigConditions {
        if err := condition.Check(&c.parseMetadata, c); err != nil {
            resultErrors = append(resultErrors, err)
        }
    }

    return resultErrors
}

type WebdavConfig struct {
    WebdavHost  string `toml:"webdab_host"`
    UseAuth     bool `toml:"use_auth"`

    WebdavAuthConfiguration *WebdavAuthetication `toml:"auth"`
}

type WebdavAuthetication struct {
    Username string `toml:"username"`
    Password string `toml:"password"`
}
