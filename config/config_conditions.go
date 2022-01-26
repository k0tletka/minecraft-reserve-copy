package config

import (
    "strings"

    "github.com/BurntSushi/toml"
)

type validCondition interface {
    Check(*toml.MetaData, *Configuration)  error
}

// Condition that checks defined metadata
type validConditionDefinedMetadata struct {
    metadataString string
}

func (dm validConditionDefinedMetadata) Check(metadata *toml.MetaData, conf *Configuration) error {
    if !metadata.IsDefined(strings.Split(dm.metadataString, ".")...) {
        return &ConfigOptionIsNotDefined{dm.metadataString}
    }

    return nil
}

// Condition type that wraps arbitrary function
type validConditionFunc func(*toml.MetaData, *ApplicationConfig) error

func (f validConditionFunc) Check(m *toml.MetaData, c *ApplicationConfig) error {
    return f(m, c)
}
