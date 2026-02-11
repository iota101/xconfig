package xconfig

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type yamlConfig struct {
	data map[string]any
}

func FromYAML(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("xconfig: read %s: %w", path, err)
	}

	var parsed map[string]any
	if err := yaml.Unmarshal(data, &parsed); err != nil {
		return nil, fmt.Errorf("xconfig: parse yaml: %w", err)
	}

	return &yamlConfig{data: parsed}, nil
}

func (c *yamlConfig) Get(key K) Value {
	val, found := c.lookup(string(key))
	if !found {
		return emptyValue(string(key))
	}
	return newValue(val, true, string(key))
}

func (c *yamlConfig) Has(key K) bool {
	_, found := c.lookup(string(key))
	return found
}

func (c *yamlConfig) lookup(key string) (any, bool) {
	parts := strings.Split(key, ".")
	var current any = c.data

	for _, part := range parts {
		m, ok := current.(map[string]any)
		if !ok {
			return nil, false
		}
		val, ok := m[part]
		if !ok {
			return nil, false
		}
		current = val
	}

	return current, true
}
