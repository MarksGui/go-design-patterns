package factory_method

import "errors"

type Config struct {
	Type string
}

type Parser interface {
	Parse() *Config
}

func NewParser(ext string) (IParserFactory, error) {
	if ext == "json" {
		return &JsonParserFactory{}, nil
	} else if ext == "xml" {
		return &XmlParserFactory{}, nil
	} else if ext == "yaml" {
		return &YamlParserFactory{}, nil
	}

	return nil, errors.New("错误的解析类型")
}

type IParserFactory interface {
	CreateParser() Parser
}

type JsonParserFactory struct{}

func (p *JsonParserFactory) CreateParser() Parser {
	return NewJsonParser()
}

type XmlParserFactory struct{}

func (p *XmlParserFactory) CreateParser() Parser {
	return NewXmlParser()
}

type YamlParserFactory struct{}

func (p *YamlParserFactory) CreateParser() Parser {
	return NewYamlParser()
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
	return &Config{Type: "xml"}
}

type YamlParser struct{}

func NewYamlParser() *YamlParser {
	return &YamlParser{}
}

func (p *YamlParser) Parse() *Config {
	return &Config{Type: "yaml"}
}
