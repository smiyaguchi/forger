package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Schema *Schema
}

type Schema struct {
	Fields []Field
}

type Field struct {
	Name   string
	Typ    string
	Schema *Schema
}

func Load(path string) (*Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	m := make(map[string]interface{})
	if err = yaml.NewDecoder(f).Decode(&m); err != nil {
		return nil, err
	}
	if _, ok := m["schema"]; !ok {
		return nil, fmt.Errorf("required schmea field")
	}

	s, err := parseSchema(m["schema"])
	if err != nil {
		return nil, err
	}
	return &Config{Schema: s}, nil
}

func parseSchema(data interface{}) (*Schema, error) {
	schema := &Schema{}
	array := data.([]interface{})
	for _, v := range array {
		switch v.(type) {
		case map[interface{}]interface{}:
			m := v.(map[interface{}]interface{})
			f := Field{}
			f.Name = m[interface{}("name")].(string)
			f.Typ = m[interface{}("type")].(string)
			if vv, ok := m[interface{}("schema")]; ok {
				s, err := parseSchema(vv)
				if err != nil {
					return nil, err
				}
				f.Schema = s
			}
			schema.Fields = append(schema.Fields, f)
		}
	}
	return schema, nil
}
