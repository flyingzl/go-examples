package filters

import (
	"errors"
)

type SumPipe struct {
}

var SumPipeParamsError = errors.New("参数必须为整形切片列表")

func (p *SumPipe) Process(data Request) (Response, error) {
	ret := 0
	if convertedSlice, ok := data.([]int); !ok {
		return 0, SumPipeParamsError
	} else {
		for _, value := range convertedSlice {
			ret += value
		}
		return ret, nil
	}
}

func NewSumPipe() *SumPipe {
	return &SumPipe{}
}
