package filters

import (
	"errors"
	"strings"
)

type StringsPipe struct {
	Delimiter string
}

var StringsPipeError = errors.New("必须为字符串")

func (p *StringsPipe) Process(data Request) (Response, error) {
	if v, ok := data.(string); !ok {
		return nil, StringsPipeError
	} else {
		return strings.Split(v, p.Delimiter), nil
	}
}

// NewStringPipe used to generate StringsPipe
func NewStringPipe(delimiter string) *StringsPipe {
	return &StringsPipe{Delimiter: delimiter}
}
