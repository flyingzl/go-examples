package filters

import (
	"errors"
	"strconv"
)

type ToIntPipe struct {
}

var ToIntPipeParamsError = errors.New("参数必须为字符串切片列表")

func (p *ToIntPipe) Process(data Request) (Response, error) {
	ret := []int{}
	if convertedSlice, ok := data.([]string); !ok {
		return nil, ToIntPipeParamsError
	} else {
		for _, value := range convertedSlice {
			value, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			ret = append(ret, value)
		}
		return ret, nil
	}
}

func NewToIntPipe() *ToIntPipe {
	return &ToIntPipe{}
}
