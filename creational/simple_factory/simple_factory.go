package simple_factory

import (
	"errors"
)

type Config struct {
	Type string
}

type Parser interface {
	Parse() *Config
}

func NewParser(ext string) (Parser, error) {
	if ext == "json" {
		return NewJsonParser(), nil
	} else if ext == "xml" {
		return NewXmlParser(), nil
	} else if ext == "yaml" {
		return NewYamlParser(), nil
	}

	return nil, errors.New("错误的解析类型")
}

type JsonParser struct{}

func NewJsonParser() *JsonParser {
	return &JsonParser{}
}

func (p *JsonParser) Parse() *Config {
	return &Config{Type: "json"}
}

type XmlParser struct{}

func NewXmlParser() *XmlParser {
	return &XmlParser{}
}

func (p *XmlParser) Parse() *Config {
	// todo
	return &Config{Type: "xml"}
}

type YamlParser struct{}

func NewYamlParser() *YamlParser {
	return &YamlParser{}
}

func (p *YamlParser) Parse() *Config {
	return &Config{Type: "yaml"}
}
